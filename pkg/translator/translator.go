package translator

import "context"

type Translator interface {
	Translate(ctx context.Context, word string) (*TranslationResponse, error)
}

type TranslationResponse struct {
	Meanings []string `json:"meanings"`
	Examples []string `json:"examples"`
	Synonyms []string `json:"synonyms"`
}
