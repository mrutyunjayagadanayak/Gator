package command

import (
	"Gator/internal/state"
	"context"
	"fmt"
)

func HandlerReset(s *state.State, cmd Command) error {
	err := s.DB.DeleteFeedFollow(context.Background())

	if err != nil {
		return fmt.Errorf("error resetting database - %v", err)
	}

	err = s.DB.Deletefeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting database - %v", err)
	}
	err = s.DB.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting database - %v", err)
	}
	fmt.Println("Database reset successfully")
	return nil
}
