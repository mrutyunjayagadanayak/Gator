package command

import (
	"Gator/internal/database"
	"Gator/internal/state"
	"context"
	"fmt"
	"strconv"
)

func HandlerBrowse(s *state.State, cmd Command, user database.User) error {
	var browseLimit int
	if len(cmd.Args) == 1 {
		var err error
		browseLimit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit provided: %v", err)
		}
	} else {
		browseLimit = 2
	}
	posts, err := s.DB.GetAllPostsByUser(context.Background(), database.GetAllPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(browseLimit),
	})
	if err != nil {
		return fmt.Errorf("error fetching feeds - %v", err)
	}
	if len(posts) == 0 {
		fmt.Println("No feeds found")
		return nil
	}
	for _, post := range posts {
		if post.Title.Valid {
			printPost(post)
		} else {
			fmt.Println("Feed: (no title)")
		}
		fmt.Println()
	}
	return nil
}

func printPost(post database.Post) {
	fmt.Println("----------------------")
	fmt.Printf("Title: %s\n", post.Title.String)
	fmt.Printf("URL: %s\n", post.Url)
	fmt.Printf("Created: %s\n", post.CreatedAt)
	if post.Description.Valid {
		fmt.Printf("Description: %s\n", post.Description.String)
	}
	if post.PublishedAt.Valid {
		fmt.Printf("Published: %s\n", post.PublishedAt.Time)
	}
}
