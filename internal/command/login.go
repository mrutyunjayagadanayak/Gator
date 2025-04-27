package command

import (
	"Gator/internal/state"

	"context"
	"fmt"
)

func HandlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no username given")
	}
	user, err := s.DB.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("user %s doesn't exist - %v", cmd.Args[0], err)
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("unable to set username - %v", err)
	}

	fmt.Printf("Username set to %s\n", cmd.Args[0])
	return nil
}
