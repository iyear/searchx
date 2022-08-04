package model

import (
	"github.com/iyear/searchx/app/bot/run/internal/i18n"
	"github.com/iyear/searchx/pkg/storage"
	"go.uber.org/zap"
)

type Scope struct {
	Storage  *Storage
	Template *i18n.Template
	Log      *zap.SugaredLogger
}

type Storage struct {
	KV     storage.KV
	Search storage.Search
	Cache  storage.Cache
}
