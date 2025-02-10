package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/hamidzmi/vocab-trainer/pkg/translator"
)

const (
	TranslationMenu  = "Enter a word (or type '/back' to return to the main menu):"
	ReturnToMainMenu = "Returning to the main menu..."
)

type Menu struct {
	Options []string
}

var Reset = "\033[0m"
var Green = "\033[32m"

func main() {
	translator := translator.NewOllamaTranslator("German", "B1")
	showMainMenu(os.Stdout, os.Stdin, translator)
}

func showMainMenu(output io.Writer, input io.Reader, translator translator.Translator) {
	mainMenu := Menu{
		[]string{
			"Learn a new word",
			"Review learnt words",
			"Quiz",
			"Exit",
		},
	}

	mainMenuString := "Let me know what do you want to do now:\n"
	for index, text := range mainMenu.Options {
		mainMenuString += strconv.Itoa(index+1) + ". " + text + "\n"
	}
	mainMenuString += "Enter your choice: "

	reader := bufio.NewReader(input)

	for {
		fmt.Fprint(output, mainMenuString)
		choice, _ := reader.ReadString('\n')
		switch choice {
		case "1\n":
			showTranslationMenu(output, reader, translator)
		case "2\n":
			fmt.Fprintln(output, "Reviewing learnt words...")
		case "3\n":
			fmt.Fprintln(output, "Get ready for a quiz...")
		case "4\n":
			fmt.Fprintln(output, "Goodbye!")
			return
		}
	}
}

func showTranslationMenu(output io.Writer, reader *bufio.Reader, translator translator.Translator) {
	for {
		fmt.Fprintln(output, TranslationMenu)

		word, _ := reader.ReadString('\n')
		word = word[:len(word)-1] // Trim newline
		if word == "/back" {
			fmt.Fprintln(output, ReturnToMainMenu)
			return
		}

		response, err := translator.Translate(context.Background(), word)
		if err != nil {
			fmt.Println(err.Error())
		}

		formatTranslationOutput(output, response)
	}
}

func formatTranslationOutput(output io.Writer, resp *translator.TranslationResponse) {
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
