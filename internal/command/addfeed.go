package command

import (
	"Gator/internal/database"
	"Gator/internal/feed"
	"Gator/internal/state"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func HandlerAddFeed(s *state.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("not enough arguments provided")
	}
	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]
	currentUser, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}
	_, err = feed.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %v", err)
	}
	rFeed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    currentUser.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %v", err)
	}
	fmt.Printf("Feed added: %v\n", rFeed)
	return nil

}
