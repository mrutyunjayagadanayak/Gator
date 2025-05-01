package internal

import (
	"Gator/internal/command"
	"Gator/internal/database"
	"Gator/internal/state"
	"context"
	"fmt"
)

func MiddlewareLoggedIN(handler func(s *state.State, cmd command.Command, user database.User) error) func(*state.State, command.Command) error {
	return func(s *state.State, cmd command.Command) error {
		user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("user %s doesn't exist - %v", s.Config.CurrentUserName, err)
		}
		return handler(s, cmd, user)
	}
}
