package command

import (
	"Gator/internal/database"
	"Gator/internal/state"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func printFeedFollow(userName string, feedName string) {
	println("Feed Follow:")
	println("User Name: " + userName)
	println("Feed Name: " + feedName)
}

func HandlerFollow(s *state.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}
	user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}
	feed, err := s.DB.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to get feed: %v", err)
	}
	ffRow, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %v", err)
	}
	fmt.Println("Feed Follow Created:")
	printFeedFollow(ffRow.UserName, ffRow.FeedName)
	return nil
}

func HandlerListFeedFollows(s *state.State, cmd Command) error {
	user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}
	feedFollows, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get feed follows: %v", err)
	}
	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found.")
		return nil
	}
	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, ff := range feedFollows {
		fmt.Printf("* %s\n", ff.FeedName)
	}
	return nil
}
