package translator

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type LLMClient interface {
	Call(ctx context.Context, prompt string, options ...llms.CallOption) (string, error)
}

type OllamaTranslator struct {
	llm    LLMClient
	prompt string
}

func NewOllamaTranslator(language string, level string) *OllamaTranslator {
	llm, _ := ollama.New(ollama.WithModel("llama3.2"))

	prompt := fmt.Sprintf(`
		You are an advanced language teacher. Your task is to translate a %s word/phrase into English and provide structured learning material.
		
		Instructions:
		- **Meanings** should be in English.
		- **Examples** should be in %s, without English translation, suitable for %s learners (clear syntax, common vocabulary, moderate sentence length).
		- **Synonyms** should be in %s.
		- The response must be **strictly valid JSON** with no introductory text or explanations
		- Use this format:
		  {
		    "meanings": ["meaning1", "meaning2", ...],
		    "examples": ["example1", "example2", ...],
		    "synonyms": ["synonym1", "synonym2", ...]
		  }
		
		Now, translate the following %s word/phrase: 
		`, language, language, level, language, language)

	return &OllamaTranslator{
		llm:    llm,
		prompt: prompt + "%s",
	}
}

func (o *OllamaTranslator) Translate(ctx context.Context, word string) (*TranslationResponse, error) {
	completePrompt := fmt.Sprintf(o.prompt, word)
	completion, err := o.llm.Call(ctx, completePrompt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate response: %v", err)
	}

	var translation TranslationResponse
	err = json.Unmarshal([]byte(completion), &translation)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v, response: %s", err, completion)
	}

	return &translation, nil
}
