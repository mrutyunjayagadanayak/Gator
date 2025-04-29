package command

import (
	"Gator/internal/state"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func HandlerGetAllFeeds(s *state.State, cmd Command) error {
	users := make(map[uuid.UUID]string)

	feeds, err := s.DB.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		userName, ok := users[feed.UserID]
		if !ok {
			user, err := s.DB.GetUserById(context.Background(), feed.UserID)
			if err != nil {
				return fmt.Errorf("failed to get user with id - %v: %v", feed.UserID, err)
			}
			users[feed.UserID] = user.Name
			userName = user.Name
		}
		fmt.Printf("Feed Name: %s, Feed URL: %s, User Name: %s\n", feed.Name, feed.Url, userName)
	}
	return nil
}
