package model

import (
	"github.com/iyear/searchx/app/usr/run/internal/i18n"
	"github.com/iyear/searchx/pkg/storage"
	"go.uber.org/zap"
)

type BotScope struct {
	Storage  *Storage
	Template *i18n.Template
	Log      *zap.SugaredLogger
}

type Storage struct {
	KV     storage.KV
	Search storage.Search
	Cache  storage.Cache
}
