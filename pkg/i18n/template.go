package i18n

import (
	"strconv"
	"strings"
	"text/template"
	"time"
)

type Text struct {
	tmpl *template.Template
}

func NewText(t string) (Text, error) {
	tmpl, err := template.New(strconv.FormatInt(time.Now().UnixNano(), 10)).Parse(t)
	if err != nil {
		return Text{}, err
	}
	return Text{tmpl: tmpl}, nil
}

func (t Text) T(data interface{}) string {
	if t.tmpl == nil {
		return ""
	}
	var buf strings.Builder
	err := t.tmpl.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
