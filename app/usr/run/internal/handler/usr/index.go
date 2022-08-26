package usr

import (
	"context"
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/app/usr/internal/index"
	"github.com/iyear/searchx/app/usr/run/internal/util"
	"github.com/iyear/searchx/pkg/keygen"
	"github.com/iyear/searchx/pkg/storage/search"
)

func Index(ctx context.Context, msg tg.MessageClass, e tg.Entities) error {
	sp := util.GetUsrScope(ctx)

	mm, ok := msg.(*tg.Message)
	if !ok {
		return nil
	}

	m, ok := index.Message(mm, e)
	if !ok {
		return nil
	}

	sp.Log.Debugw("new message", "chatID", m.Chat, "chatType", m.ChatType, "chatName", m.ChatName, "msgID", m.ID, "senderID", m.Sender, "senderName", m.SenderName, "text", m.Text, "date", m.Date)

	data, err := m.Encode()
	if err != nil {
		return err
	}
	return sp.Storage.Search.Index(ctx, []*search.Item{{
		ID:   keygen.SearchMsgID(m.Chat, m.ID),
		Data: data,
	}})
}
