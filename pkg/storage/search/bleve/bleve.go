package bleve

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/custom"
	"github.com/blevesearch/bleve/v2/analysis/lang/en"
	"github.com/blevesearch/bleve/v2/analysis/token/lowercase"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/creasty/defaults"
	"github.com/iyear/searchx/pkg/utils"
	"github.com/iyear/searchx/pkg/validator"
	"github.com/mitchellh/mapstructure"
)

type Bleve struct {
	index bleve.Index
}

type Options struct {
	Path string `mapstructure:"path" default:"data/index"`
	Dict string `mapstructure:"dict" validate:"file" default:"data/dict.txt"`
}

const (
	jieba = "jieba"
)

func New(options map[string]interface{}) (*Bleve, error) {
	// add jieba to registry
	registry.RegisterTokenizer(jieba, JiebaTokenizerConstructor)

	var ops Options

	if err := mapstructure.WeakDecode(options, &ops); err != nil {
		return nil, err
	}

	if err := defaults.Set(&ops); err != nil {
		return nil, err
	}

	if err := validator.Struct(&ops); err != nil {
		return nil, err
	}

	mapping := bleve.NewIndexMapping()
	err := mapping.AddCustomTokenizer(jieba, map[string]interface{}{
		"file": ops.Dict,
		"type": jieba,
	})
	if err != nil {
		return nil, err
	}

	err = mapping.AddCustomAnalyzer(jieba, map[string]interface{}{
		"type":          custom.Name,
		"tokenizer":     jieba,
		"token_filters": []string{en.PossessiveName, lowercase.Name, en.StopName},
	})
	if err != nil {
		return nil, err
	}

	mapping.DefaultAnalyzer = jieba
	mapping.DefaultMapping.StructTagKey = "index"

	var index bleve.Index

	if !utils.FS.PathExist(ops.Path) {
		index, err = bleve.New(ops.Path, mapping)
		if err != nil {
			return nil, err
		}
		if err = index.Close(); err != nil {
			return nil, err
		}
	}

	if index, err = bleve.Open(ops.Path); err != nil {
		return nil, err
	}

	return &Bleve{index: index}, nil
}
