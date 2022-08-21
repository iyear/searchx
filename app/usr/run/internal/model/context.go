package model

import (
	"github.com/gotd/td/telegram"
	"github.com/iyear/searchx/app/usr/run/internal/i18n"
	"github.com/iyear/searchx/pkg/storage"
	"go.uber.org/zap"
)

type BotScope struct {
	Storage  *storage.Storage
	Template *i18n.Template
	Log      *zap.SugaredLogger
}

type UsrScope struct {
	Storage *storage.Storage
	Log     *zap.SugaredLogger
	Client  *telegram.Client
}
