package translator

import (
	"context"
	"testing"
)

func TestOllamaTranslator_Translate_Success(t *testing.T) {
	mockJSON := `{
		"meanings": ["dog"],
		"examples": ["Der Hund bellt.", "Mein Hund ist freundlich."],
		"synonyms": ["Welpe", "Hündchen"]
	}`

	mockTranslator := NewMockOllamaTranslator(mockJSON, false, t)

	resp, err := mockTranslator.Translate(context.Background(), "Hund")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(resp.Meanings) == 0 || resp.Meanings[0] != "dog" {
		t.Errorf("expected meaning 'dog', got %v", resp.Meanings)
	}
}

func TestOllamaTranslator_Translate_ErrorHandling(t *testing.T) {
	mockTranslator := NewMockOllamaTranslator("", true, t)

	_, err := mockTranslator.Translate(context.Background(), "Hund")
	if err == nil {
		t.Fatalf("expected an error but got nil")
	}
}

func TestOllamaTranslator_Translate_VerifyPrompt(t *testing.T) {
	mockJSON := `{
		"meanings": ["dog"],
		"examples": ["Der Hund bellt.", "Mein Hund ist freundlich."],
		"synonyms": ["Welpe", "Hündchen"]
	}`

	mockLLM := &MockLLM{MockResponse: mockJSON}
	translator := &OllamaTranslator{llm: mockLLM, prompt: "Translate: %s"}

	_, err := translator.Translate(context.Background(), "Hund")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(mockLLM.ReceivedPrompts) == 0 {
		t.Fatalf("expected at least one prompt but got none")
	}
}
