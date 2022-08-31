package i18n

import (
	"github.com/mitchellh/mapstructure"
	tele "gopkg.in/telebot.v3"
	"hash/crc64"
	"reflect"
	"strconv"
)

func TextHook() mapstructure.DecodeHookFunc {
	return func(f reflect.Value, t reflect.Value) (interface{}, error) {
		if f.Kind() == reflect.String && t.Type() == reflect.TypeOf(Text{}) {
			return NewText(f.String())
		}
		return f.Interface(), nil
	}
}

func InlineButtonHook() mapstructure.DecodeHookFunc {
	return func(f reflect.Value, t reflect.Value) (interface{}, error) {
		if f.Kind() == reflect.String && t.Type() == reflect.TypeOf(tele.InlineButton{}) {
			return tele.InlineButton{Unique: btnUnique(f.String()), Text: f.String()}, nil
		}
		return f.Interface(), nil
	}
}

func btnUnique(text string) string {
	return strconv.FormatUint(crc64.Checksum([]byte(string([]rune(text)[0])), crc64.MakeTable(crc64.ISO)), 10)
}
