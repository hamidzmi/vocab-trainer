package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const Menu = "Let me know what do you want to do now:\n" +
	"1. Learn a new word\n" +
	"2. Review learnt words\n" +
	"3. Quiz\n" +
	"4. Exit\n" +
	"Enter your choice: "

var Reset = "\033[0m"
var Green = "\033[32m"

func MainMenu(output io.Writer, input io.Reader) {
	reader := bufio.NewReader(input)

	for {
		fmt.Fprint(output, Menu)
		choice, _ := reader.ReadString('\n')
		switch choice {
		case "1\n":
			translationMenu(output, reader)
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

func translationMenu(output io.Writer, reader *bufio.Reader) {
	for {
		fmt.Fprintln(output, "Enter a word (or type '/back' to return to the main menu):")

		word, _ := reader.ReadString('\n')
		word = word[:len(word)-1] // Trim newline
		if word == "/back" {
			fmt.Fprintln(output, "Returning to the main menu...")
			return
		}

		response, err := Translate(word)
		if err != nil {
			fmt.Println(err.Error())
		}

		FormatTranslationOutput(output, response)
	}
}

func main() {
	MainMenu(os.Stdout, os.Stdin)
}
