package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMainMenuOutput(t *testing.T) {
	t.Run("select 1 from main menu", func(t *testing.T) {
		input := strings.NewReader("1\n4\n")
		buffer := bytes.Buffer{}

		MainMenu(&buffer, input)

		got := buffer.String()
		want := Menu + "Adding new word...\n" + Menu + "Goodbye!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("select 2 from main menu", func(t *testing.T) {
		input := strings.NewReader("2\n4\n")
		buffer := bytes.Buffer{}

		MainMenu(&buffer, input)

		got := buffer.String()
		want := Menu +
			"Reviewing learnt words...\n" +
			Menu +
			"Goodbye!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("select 3 from main menu", func(t *testing.T) {
		input := strings.NewReader("3\n4\n")
		buffer := bytes.Buffer{}

		MainMenu(&buffer, input)

		got := buffer.String()
		want := Menu +
			"Get ready for a quiz...\n" +
			Menu +
			"Goodbye!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("select 4 from main menu", func(t *testing.T) {
		input := strings.NewReader("4\n")
		buffer := bytes.Buffer{}

		MainMenu(&buffer, input)

		got := buffer.String()
		want := Menu + "Goodbye!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
