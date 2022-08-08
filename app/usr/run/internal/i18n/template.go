package i18n

type Template struct {
	Text   *TemplateText   `mapstructure:"text"`
	Button *TemplateButton `mapstructure:"button"`
}

type TemplateText struct {
}

type TemplateButton struct {
}
