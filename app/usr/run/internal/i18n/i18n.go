package i18n

import (
	"github.com/iyear/searchx/pkg/i18n"
	"github.com/iyear/searchx/pkg/utils"
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
		if err = i18n.Read(f, &t, i18n.InlineButtonHook()); err != nil {
			return err
		}
		lang := utils.FS.GetFileName(f)
		langs = append(langs, lang)
		m[lang] = &t
	}

	Templates = m

	return nil
}
