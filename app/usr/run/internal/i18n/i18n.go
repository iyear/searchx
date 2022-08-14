package i18n

import (
	"github.com/iyear/searchx/pkg/i18n"
	"github.com/iyear/searchx/pkg/utils"
	"github.com/mitchellh/mapstructure"
	tele "gopkg.in/telebot.v3"
	"hash/crc64"
	"reflect"
	"strconv"
)

var Templates = make(map[string]*Template)

func Init(dir string) error {
	files, err := i18n.Walk(dir)
	if err != nil {
		return err
	}

	m := make(map[string]*Template)
	langs := make([]string, 0, len(files))
	for _, f := range files {
		t := Template{}
		if err = i18n.Read(f, &t, readHook()); err != nil {
			return err
		}
		lang := utils.FS.GetFileName(f)
		langs = append(langs, lang)
		m[lang] = &t
	}

	Templates = m

	return nil
}

func readHook() mapstructure.DecodeHookFunc {
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
