package command

import (
	"Gator/internal/state"
	"fmt"
)

func HandlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("No username given")
	}
	err := s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Unable to set Username - %v", err)
	}

	fmt.Printf("Username set to %s\n", cmd.Args[0])
	return nil
}
