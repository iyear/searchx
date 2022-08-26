package source

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/telegram/message/peer"
	"github.com/gotd/td/telegram/query"
	"github.com/gotd/td/telegram/query/messages"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	"github.com/iyear/searchx/app/usr/internal/config"
	"github.com/iyear/searchx/app/usr/internal/index"
	"github.com/iyear/searchx/app/usr/internal/sto"
	"github.com/iyear/searchx/pkg/keygen"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/storage/search"
	"github.com/iyear/searchx/pkg/utils"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"time"
)

func Start(ctx context.Context, cfg string, date int) error {
	if err := config.Init(cfg); err != nil {
		return err
	}
	color.Blue("Config loaded")

	_search, kv, _, err := storage.Init(config.C.Storage)
	if err != nil {
		return err
	}
	color.Blue("Storage initialized")

	dialer, err := utils.ProxyFromURL(config.C.Proxy)
	if err != nil {
		return err
	}

	c := telegram.NewClient(config.C.Account.ID, config.C.Account.Hash, telegram.Options{
		Resolver: dcs.Plain(dcs.PlainOptions{
			Dial: dialer.DialContext,
		}),
		SessionStorage: sto.NewSession(kv, false),
		Logger:         zap.NewNop(),
	})

	return c.Run(ctx, func(ctx context.Context) error {
		status, err := c.Auth().Status(ctx)
		if err != nil {
			return err
		}
		if !status.Authorized {
			return fmt.Errorf("not authorized. please login first")
		}

		color.Blue("Authorized: %s", status.User.Username)

		color.Blue("Get Blocked Dialogs...")
		blocks, err := query.GetBlocked(c.API()).BatchSize(100).Collect(ctx)
		if err != nil {
			return err
		}

		blockids := make(map[int64]struct{})
		for _, b := range blocks {
			blockids[GetPeerID(b.Contact.PeerID)] = struct{}{}
		}

		color.Blue("Get All Dialogs...")
		dialogs, err := query.GetDialogs(c.API()).BatchSize(100).Collect(ctx)
		if err != nil {
			return err
		}

		color.Blue("Indexing...")
		wg, errctx := errgroup.WithContext(ctx)
		wg.SetLimit(2)

		for _, d := range dialogs {
			d := d
			// fmt.Printf("id: %d, name: %s\n", GetInputPeerID(d.Peer), GetInputPeerName(d.Peer, d.Entities))
			if _, blocked := blockids[GetInputPeerID(d.Peer)]; blocked {
				continue
			}

			time.Sleep(time.Second)
			wg.Go(func() error {
				return fetch(errctx, _search, d.Peer, d.Entities, query.Messages(c.API()).GetHistory(d.Peer), date)
			})
		}

		return wg.Wait()
	})
}

func fetch(ctx context.Context, _search storage.Search, peer tg.InputPeerClass, e peer.Entities, builder *messages.GetHistoryQueryBuilder, date int) error {
	id := GetInputPeerID(peer)

	batchSize := 100
	indexSize := 20

start:
	start := time.Now()
	iter := builder.BatchSize(batchSize).Iter()
	count := 0
	msgs := make([]*search.Item, 0, indexSize)

	for {
		if !iter.Next(ctx) {
			if iter.Err() == nil {
				break
			}

			if dur, ok := tgerr.AsFloodWait(iter.Err()); ok {
				time.Sleep(dur + time.Second)
				goto start
			}

			return iter.Err()
		}

		msg := iter.Value()
		if msg.Msg.GetDate() < date {
			fmt.Printf("id: %d,name: %s count: %d,took: %s\n", id, GetInputPeerName(peer, e), count, time.Since(start))
			break
		}

		// index msg
		m, ok := msg.Msg.(*tg.Message)
		if !ok {
			continue
		}

		data, ok := index.Message(m, tg.Entities{
			Short:    false,
			Users:    e.Users(),
			Chats:    e.Chats(),
			Channels: e.Channels(),
		})
		if !ok {
			continue
		}

		dd, err := data.Encode()
		if err != nil {
			return err
		}

		msgs = append(msgs, &search.Item{
			ID:   keygen.SearchMsgID(data.Chat, data.ID),
			Data: dd,
		})
		count++

		if len(msgs) == indexSize {
			if err = _search.Index(ctx, msgs); err != nil {
				return err
			}
			msgs = make([]*search.Item, 0, indexSize)
		}

		if count%batchSize == 0 {
			time.Sleep(700 * time.Millisecond)
		}
	}

	if len(msgs) > 0 {
		if err := _search.Index(ctx, msgs); err != nil {
			return err
		}
	}

	return nil
}
