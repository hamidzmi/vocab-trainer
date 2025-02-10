package main

import (
	"bytes"
	"strings"
	"testing"
)

const MainMenu = "Let me know what do you want to do now:\n" +
	"1. Learn a new word\n" +
	"2. Review learnt words\n" +
	"3. Quiz\n" +
	"4. Exit\n" +
	"Enter your choice: "

func TestMainMenuOutput(t *testing.T) {
	mockTranslator := &MockTranslator{}

	t.Run("select 1 from main menu", func(t *testing.T) {
		input := strings.NewReader("1\n/back\n4\n")
		buffer := bytes.Buffer{}

		showMainMenu(&buffer, input, mockTranslator)

		got := buffer.String()
		want := MainMenu +
			TranslationMenu + "\n" +
			ReturnToMainMenu + "\n" +
			MainMenu +
			"Goodbye!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("select 2 from main menu", func(t *testing.T) {
		input := strings.NewReader("2\n4\n")
		buffer := bytes.Buffer{}

		showMainMenu(&buffer, input, mockTranslator)

		got := buffer.String()
		want := MainMenu +
			"Reviewing learnt words...\n" +
			MainMenu +
			"Goodbye!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("select 3 from main menu", func(t *testing.T) {
		input := strings.NewReader("3\n4\n")
		buffer := bytes.Buffer{}

		showMainMenu(&buffer, input, mockTranslator)

		got := buffer.String()
		want := MainMenu +
			"Get ready for a quiz...\n" +
			MainMenu +
			"Goodbye!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("select 4 from main menu", func(t *testing.T) {
		input := strings.NewReader("4\n")
		buffer := bytes.Buffer{}

		showMainMenu(&buffer, input, mockTranslator)

		got := buffer.String()
		want := MainMenu + "Goodbye!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
