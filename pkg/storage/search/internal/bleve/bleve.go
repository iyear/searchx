package bleve

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/custom"
	"github.com/blevesearch/bleve/v2/analysis/lang/en"
	"github.com/blevesearch/bleve/v2/analysis/token/lowercase"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/iyear/searchx/pkg/utils"
	"github.com/mitchellh/mapstructure"
	"path"
)

type Bleve struct {
	index bleve.Index
}

type Options struct {
	Path string `mapstructure:"path"`
	Dict string `mapstructure:"dict"`
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

	var index bleve.Index

	indexPath := path.Join(ops.Path)
	if !utils.PathExist(indexPath) {
		index, err = bleve.New(indexPath, mapping)
		if err != nil {
			return nil, err
		}
		if err = index.Close(); err != nil {
			return nil, err
		}
	}

	if index, err = bleve.Open(indexPath); err != nil {
		return nil, err
	}

	return &Bleve{index: index}, nil
}
