package main

import (
	"fmt"

	"github.com/aashaybelekar/go-do/cmd"
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
                                                                      
`
	fmt.Println(app_logo)
	cmd.Execute()

}
