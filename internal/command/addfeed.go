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

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", user.Name)
}

func HandlerAddFeed(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("not enough arguments provided")
	}
	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	//currentUser, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
	//if err != nil {
	//	return err
	//}
	_, err := feed.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %v", err)
	}
	rFeed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %v", err)
	}

	feedFollow, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    rFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %v", err)
	}

	fmt.Printf("Feed added: %v\n", rFeed)
	printFeed(rFeed, user)
	fmt.Println()
	fmt.Println("Feed followed successfully:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("=====================================")
	return nil

}
