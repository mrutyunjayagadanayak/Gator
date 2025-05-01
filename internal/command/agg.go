package command

import (
	"Gator/internal/database"
	"Gator/internal/feed"
	"Gator/internal/state"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func HandlerAgg(s *state.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: gator agg [time_between_reqs]")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid time duration: %v", err)
	}
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			fmt.Printf("Error scraping feeds: %v\n", err)
		}
	}
	//feed, err := feed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	//if err != nil {
	//	return fmt.Errorf("failed to fetch feed: %v", err)
	//}
	//fmt.Printf("%+v\n", feed)
	return nil
}

func scrapeFeeds(s *state.State) error {
	feeed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.DB.UpdateFeedFetchTime(context.Background(), database.UpdateFeedFetchTimeParams{
		ID:            feeed.ID,
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return err
	}
	feedData, err := feed.FetchFeed(context.Background(), feeed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %v", err)
	}
	for _, item := range feedData.Channel.Item {
		fmt.Println(item.Title)
	}
	return nil
}
