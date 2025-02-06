package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/tmc/langchaingo/llms/ollama"
)

const model = "llama3.2"
const prompt = `
	You are an advanced language teacher. Your task is to translate a German word/phrase into English and provide structured learning material.
	
	Instructions:
	- **Meanings** should be in English.
	- **Examples** should be in German, without English translation, suitable for B1 learners (clear syntax, common vocabulary, moderate sentence length).
	- **Synonyms** should be in German.
	- The response must be **strictly valid JSON** with no introductory text or explanations
	- Use this format:
	  {
	    "meanings": ["meaning1", "meaning2", ...],
	    "examples": ["example1", "example2", ...],
	    "synonyms": ["synonym1", "synonym2", ...]
	  }
	
	Now, translate the following German word/phrase: "%s".
	`

type TranslationResponse struct {
	Meanings []string `json:"meanings"`
	Examples []string `json:"examples"`
	Synonyms []string `json:"synonyms"`
}

func Translate(word string) (*TranslationResponse, error) {
	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	completePrompt := fmt.Sprintf(prompt, word)
	completion, err := llm.Call(ctx, completePrompt)
	if err != nil {
		log.Fatal(err)
	}

	var translation TranslationResponse
	err = json.Unmarshal([]byte(completion), &translation)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v, response: %s", err, completion)
	}

	return &translation, nil
}

func FormatTranslationOutput(output io.Writer, resp *TranslationResponse) {
	fmt.Fprintln(output, "\n"+Green+"Meanings:"+Reset)
	for _, meaning := range resp.Meanings {
		fmt.Fprintf(output, "* %s\n", meaning)
	}

	fmt.Fprintln(output, "\n"+Green+"Examples:"+Reset)
	for _, example := range resp.Examples {
		fmt.Fprintf(output, "* %s\n", example)
	}

	fmt.Fprintln(output, "\n"+Green+"Synonyms:"+Reset)
	for _, synonym := range resp.Synonyms {
		fmt.Fprintf(output, "* %s\n", synonym)
	}
	fmt.Fprintf(output, "\n*********************************\n")
}
