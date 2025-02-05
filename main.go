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

func MainMenu(output io.Writer, input io.Reader) {
	reader := bufio.NewReader(input)
	for {
		fmt.Fprint(output, Menu)

		choice, _ := reader.ReadString('\n')
		switch choice {
		case "1\n":
			fmt.Fprintln(output, "Adding new word...")
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

func main() {
	MainMenu(os.Stdout, os.Stdin)
}
