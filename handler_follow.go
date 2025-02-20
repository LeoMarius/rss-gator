package main

import (
	"context"
	"fmt"

	"github.com/LeoMarius/rss-gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get the user: %w", err)
	}

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get the feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed_follows: %w", err)
	}

	fmt.Printf("* Feed Follow ok !\n")
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* UserName:      %s\n", user.Name)

	return nil
}

func handlerFollowing(s *state, cmd command) error {

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get the user: %w", err)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get the feed_follows: %w", err)
	}

	for _, follow := range follows {

		fmt.Printf("* Name:          %s\n", follow.Name)
	}
	return nil

}
