package command

import (
	"Gator/internal/state"
	"context"
	"fmt"
)

func HandlerUsers(s *state.State, cmd Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		fmt.Println("Error getting users - ", err)
	}
	for _, user := range users {
		if s.Config.CurrentUserName == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}
