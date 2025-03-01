package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aashaybelekar/go-do/cmd"
	"github.com/mattn/go-shellwords"
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

		args, err := shellwords.Parse(input)
		if err != nil {
			fmt.Println("Error parsing input:", err)
			continue
		}

		if len(args) == 0 {
			continue
		}

		os.Args = append([]string{"dummy"}, args...)

		// Execute the command using Cobra
		cmd.Execute()
	}
}
