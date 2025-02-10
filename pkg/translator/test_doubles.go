package translator

import (
	"context"
	"errors"
	"testing"

	"github.com/tmc/langchaingo/llms"
)

type MockLLM struct {
	MockResponse    string
	ShouldError     bool
	ReceivedPrompts []string
	t               *testing.T
}

func (m *MockLLM) Call(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	m.ReceivedPrompts = append(m.ReceivedPrompts, prompt)

	if m.ShouldError {
		return "", errors.New("mock error: failed to generate response")
	}

	return m.MockResponse, nil
}

func NewMockOllamaTranslator(mockResponse string, shouldError bool, t *testing.T) *OllamaTranslator {
	return &OllamaTranslator{
		llm: &MockLLM{
			MockResponse: mockResponse,
			ShouldError:  shouldError,
			t:            t,
		},
		prompt: "test prompt for translation of word %s",
	}
}
