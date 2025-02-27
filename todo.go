package main

import (
	"fmt"

	"github.com/aashaybelekar/go-do/internal/repl"
)

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
	repl.Repl()
}
