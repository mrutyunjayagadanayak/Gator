package command

import (
	"Gator/internal/database"
	"Gator/internal/state"
	"context"
	"fmt"
)

func HandlerUnfollow(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no url was given")
	}

	url := cmd.Args[0]

	err := s.DB.DeleteFeedByUserAndURL(context.Background(), database.DeleteFeedByUserAndURLParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Unfollowed %s\n", url)

	return nil
}
