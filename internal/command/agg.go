package command

import (
	"Gator/internal/feed"
	"Gator/internal/state"
	"context"
	"fmt"
)

func HandlerAgg(s *state.State, cmd Command) error {
	feed, err := feed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %v", err)
	}
	fmt.Printf("%+v\n", feed)
	return nil
}
