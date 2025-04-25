package main

import (
	"Gator/internal/config"
	"fmt"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config - %v", err)
	}

	conf.SetUser("lane")
	new_config, err := config.Read()

	fmt.Println(new_config)

}
