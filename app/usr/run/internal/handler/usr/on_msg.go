package usr

import (
	"context"
	"github.com/gotd/td/tg"
)

func OnNewMessage(ctx context.Context, e tg.Entities, update *tg.UpdateNewMessage) error {
	return Index(ctx, update.Message, e)
}

func OnEditMessage(ctx context.Context, e tg.Entities, update *tg.UpdateEditMessage) error {
	return Index(ctx, update.Message, e)
}

func OnNewScheduledMessage(ctx context.Context, e tg.Entities, update *tg.UpdateNewScheduledMessage) error {
	return Index(ctx, update.Message, e)
}

func OnNewChannelMessage(ctx context.Context, e tg.Entities, update *tg.UpdateNewChannelMessage) error {
	return Index(ctx, update.Message, e)
}

func OnEditChannelMessage(ctx context.Context, e tg.Entities, update *tg.UpdateEditChannelMessage) error {
	return Index(ctx, update.Message, e)
}
