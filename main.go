package main

import (
	"Gator/internal/command"
	"Gator/internal/config"
	"Gator/internal/state"
	"fmt"
	"os"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config - %v", err)
	}
	state := state.New(&conf)
	var commands command.Commands
	commands.Register("login", command.HandlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Error: Not enough arguments provided")
		os.Exit(1)
	}

	cmd := command.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := commands.Run(state, cmd); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

}
