package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/aashaybelekar/go-do/cmd"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	log.Print(runtime.GOOS)
	if ok { //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n> ") // Display prompt
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" || input == "quit" || input == "q" {
			fmt.Println("Goodbye!")
			break
		} else if input == "clear" {
			CallClear()
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

func main() {
	app_logo := `
                                
                                ██╗   ██╗ █████╗ ████████╗ █████╗    
                                ╚██╗ ██╔╝██╔══██╗╚══██╔══╝██╔══██╗   
                                 ╚████╔╝ ███████║   ██║   ███████║   
                                  ╚██╔╝  ██╔══██║   ██║   ██╔══██║   
                                   ██║██╗██║  ██║██╗██║██╗██║  ██║██╗
                                   ╚═╝╚═╝╚═╝  ╚═╝╚═╝╚═╝╚═╝╚═╝  ╚═╝╚═╝                              
                                      (Yet Another To-Do App)   

'exit'/'quit' to quit.
                                                                      
`
	fmt.Println(app_logo)
	repl()
}
