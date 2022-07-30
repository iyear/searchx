// Package bleve fork from github.com/wangbin/jiebago and update bleve version
package bleve

import (
	"fmt"
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/wangbin/jiebago"
	"regexp"
	"strconv"
)

var ideographRegexp = regexp.MustCompile(`\p{Han}+`)

// JiebaTokenizer is the beleve tokenizer for jiebago.
type JiebaTokenizer struct {
	seg             jiebago.Segmenter
	hmm, searchMode bool
}

func NewJiebaTokenizer(dictFilePath string, hmm, searchMode bool) (analysis.Tokenizer, error) {
	var seg jiebago.Segmenter
	err := seg.LoadDictionary(dictFilePath)
	return &JiebaTokenizer{
		seg:        seg,
		hmm:        hmm,
		searchMode: searchMode,
	}, err
}

// Tokenize cuts input into bleve token stream.
func (jt *JiebaTokenizer) Tokenize(input []byte) analysis.TokenStream {
	rv := make(analysis.TokenStream, 0)
	runeStart := 0
	start := 0
	end := 0
	pos := 1
	var width int
	var gram string
	for word := range jt.seg.Cut(string(input), jt.hmm) {
		if jt.searchMode {
			runes := []rune(word)
			width = len(runes)
			for _, step := range [2]int{2, 3} {
				if width <= step {
					continue
				}
				for i := 0; i < width-step+1; i++ {
					gram = string(runes[i : i+step])
					gramLen := len(gram)
					if frequency, ok := jt.seg.Frequency(gram); ok && frequency > 0 {
						gramStart := start + len(string(runes[:i]))
						token := analysis.Token{
							Term:     []byte(gram),
							Start:    gramStart,
							End:      gramStart + gramLen,
							Position: pos,
							Type:     detectTokenType(gram),
						}
						rv = append(rv, &token)
						pos++
					}
				}
			}
		}
		end = start + len(word)
		token := analysis.Token{
			Term:     []byte(word),
			Start:    start,
			End:      end,
			Position: pos,
			Type:     detectTokenType(word),
		}
		rv = append(rv, &token)
		pos++
		runeStart += width
		start = end
	}
	return rv
}

/*
JiebaTokenizerConstructor creates a JiebaTokenizer.

Parameter config should contains at least one parameter:

    file: the path of the dictionary file.

    hmm: optional, specify whether to use Hidden Markov Model, see NewJiebaTokenizer for details.

    search: optional, speficy whether to use search mode, see NewJiebaTokenizer for details.
*/
func JiebaTokenizerConstructor(config map[string]interface{}, _ *registry.Cache) (analysis.Tokenizer, error) {
	dictFilePath, ok := config["file"].(string)
	if !ok {
		return nil, fmt.Errorf("must specify dictionary file path")
	}
	hmm, ok := config["hmm"].(bool)
	if !ok {
		hmm = true
	}
	searchMode, ok := config["search"].(bool)
	if !ok {
		searchMode = true
	}

	return NewJiebaTokenizer(dictFilePath, hmm, searchMode)
}

func detectTokenType(term string) analysis.TokenType {
	if ideographRegexp.MatchString(term) {
		return analysis.Ideographic
	}
	_, err := strconv.ParseFloat(term, 64)
	if err == nil {
		return analysis.Numeric
	}
	return analysis.AlphaNumeric
}
