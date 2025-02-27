package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aashaybelekar/go-do/cmd"
)

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n> ") // Display prompt
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" || input == "quit" || input == "q" {
			fmt.Println("Goodbye!")
			break
		} else if input == "clear" {
			cmd.CallClear()
			continue
		}

		args := strings.Fields(input) // Split input into words
		if len(args) == 0 {
			continue
		}

		os.Args = append([]string{"dummy"}, args...)

		// Execute the command using Cobra
		cmd.Execute()
	}
}
