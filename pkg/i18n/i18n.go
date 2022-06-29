package i18n

import (
	"fmt"
	iso6391 "github.com/iyear/iso-639-1"
	"github.com/iyear/searchx/pkg/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"io/fs"
	"path/filepath"
	"reflect"
)

func Read(path string, tmpl interface{}, hook mapstructure.DecodeHookFunc) error {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(tmpl, viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(textHook(), hook))); err != nil {
		return err
	}
	return nil
}

func textHook() mapstructure.DecodeHookFunc {
	return func(f reflect.Value, t reflect.Value) (interface{}, error) {
		if f.Kind() == reflect.String && t.Type() == reflect.TypeOf(Text{}) {
			return NewText(f.String())
		}
		return f.Interface(), nil
	}
}

func Walk(dir string) ([]string, error) {
	paths := make([]string, 0)
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("unable to walk template path: %s, err:%v", path, err)
		}
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		name := utils.GetFileName(path)
		if ext != ".toml" || !iso6391.ValidCode(name) {
			return fmt.Errorf("invalid template file: %s.Please check extension or name of the file", path)
		}
		paths = append(paths, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}
