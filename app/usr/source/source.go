package source

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/gotd/contrib/middleware/floodwait"
	"github.com/gotd/contrib/middleware/ratelimit"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/telegram/message/peer"
	"github.com/gotd/td/telegram/query"
	"github.com/gotd/td/telegram/query/messages"
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/app/usr/internal/config"
	"github.com/iyear/searchx/app/usr/internal/index"
	"github.com/iyear/searchx/app/usr/internal/sto"
	"github.com/iyear/searchx/pkg/keygen"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/storage/search"
	"github.com/iyear/searchx/pkg/utils"
	"github.com/jedib0t/go-pretty/v6/progress"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"golang.org/x/time/rate"
	"time"
)

const (
	rateInterval   = 475 * time.Millisecond
	rateBucket     = 2
	rateMax        = "10k/min"
	waitGroupLimit = 2
)

func Start(ctx context.Context, cfg string, from int, to int) error {
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

	c := telegram.NewClient(config.C.App.ID, config.C.App.Hash, telegram.Options{
		Resolver: dcs.Plain(dcs.PlainOptions{
			Dial: dialer.DialContext,
		}),
		SessionStorage: sto.NewSession(kv, false),
		Logger:         zap.NewNop(),
		Middlewares: []telegram.Middleware{
			floodwait.NewSimpleWaiter(),

			// dur: 1min, wg limit:2, group: 1492447836,1354090939
			// data1 data2 every b
			// 5.9k  4.89k 475   2
			// 5.7k  4.73k 500   2
			// 5.9k  4.9k  400   2
			// 5.9k  4.9k  450   2
			// 5.9k  4.9k  400   3
			// rate limit is about 10k/min
			ratelimit.New(rate.Every(rateInterval), rateBucket),
		},
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
		blockids, err := getBlockedDialogs(ctx, c.API())
		if err != nil {
			return err
		}

		color.Blue("Get All Dialogs...")
		dialogs, err := query.GetDialogs(c.API()).BatchSize(100).Collect(ctx)
		if err != nil {
			return err
		}

		color.Blue("Indexing... %s ~ %s", time.Unix(int64(from), 0).Format("2006-01-02 15:04:05"), time.Unix(int64(to), 0).Format("2006-01-02 15:04:05"))
		color.Cyan("Fetch speed is about %s, determined by Telegram server", rateMax)

		wg, errctx := errgroup.WithContext(ctx)
		wg.SetLimit(waitGroupLimit)

		total, start := atomic.NewUint64(0), time.Now()

		// render progress
		pw := getProgress()
		go pw.Render()
		defer pw.Stop()

		for _, d := range dialogs {
			if _, blocked := blockids[utils.Telegram.GetInputPeerID(d.Peer)]; blocked {
				continue
			}
			if d.Deleted() {
				continue
			}

			d := d
			wg.Go(func() error {
				count, err := fetch(errctx, _search, pw, d.Peer, d.Entities, query.Messages(c.API()).GetHistory(d.Peer), from, to)
				if err != nil {
					return err
				}

				total.Add(uint64(count))

				return nil
			})
		}

		if err = wg.Wait(); err != nil {
			return err
		}

		time.Sleep(time.Second) // wait for progress to render the latest status
		color.Blue("Total: %d, Time: %s", total.Load(), time.Since(start))
		return nil
	})
}

func fetch(ctx context.Context, _search storage.Search, pw progress.Writer,
	peer tg.InputPeerClass, e peer.Entities, builder *messages.GetHistoryQueryBuilder,
	from int, to int) (int64, error) {
	id := utils.Telegram.GetInputPeerID(peer)
	name := utils.Telegram.GetInputPeerName(peer, e)

	batchSize := 100
	iter := builder.OffsetDate(to).BatchSize(batchSize).Iter()
	count := int64(0)
	msgs := make([]*search.Item, 0, batchSize)
	tracker := appendTracker(pw, fmt.Sprintf("%s (%d)", name, id))

	for iter.Next(ctx) {
		msg := iter.Value()
		if msg.Msg.GetDate() < from {
			break
		}

		m, ok := msg.Msg.(*tg.Message)
		if !ok {
			continue
		}
		// TODO(iyear): `e` don't have full entities, need to fetch them from telegram
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
			return 0, err
		}

		msgs = append(msgs, &search.Item{
			ID:   keygen.SearchMsgID(data.Chat, data.ID),
			Data: dd,
		})
		count++
		tracker.SetValue(count)

		if count%int64(batchSize) == 0 {
			if err = _search.Index(ctx, msgs); err != nil {
				return 0, err
			}
			msgs = make([]*search.Item, 0, batchSize)
		}
	}

	if len(msgs) > 0 {
		if err := _search.Index(ctx, msgs); err != nil {
			return 0, err
		}
	}

	tracker.MarkAsDone()

	return count, iter.Err()
}
