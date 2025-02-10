package main

import (
	"context"

	"github.com/hamidzmi/vocab-trainer/pkg/translator"
)

// MockTranslator implements translator.Translator
type MockTranslator struct{}

// Translate returns a fixed mock response
func (m *MockTranslator) Translate(ctx context.Context, word string) (*translator.TranslationResponse, error) {
	return &translator.TranslationResponse{
		Meanings: []string{"mock-meaning"},
		Examples: []string{"mock-example"},
		Synonyms: []string{"mock-synonym"},
	}, nil
}
