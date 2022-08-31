package i18n

import (
	"fmt"
	iso6391 "github.com/iyear/iso-639-1"
	"github.com/iyear/searchx/pkg/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"io/fs"
	"path/filepath"
)

func Read(path string, tmpl interface{}, hook mapstructure.DecodeHookFunc) error {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(tmpl, viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(TextHook(), hook))); err != nil {
		return err
	}
	return nil
}

func Walk(dir string) ([]string, error) {
	paths := make([]string, 0)
	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("unable to walk template path: %s, err:%v", path, err)
		}
		if entry.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		name := utils.FS.GetFileName(path)
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
