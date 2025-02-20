package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("pas un temps: %w", err)
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) error {

	fmt.Printf("c'est parti ! \n")

	nextFeed, err := s.db.GetNextFeedToFech(s.ctx)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("next feed %v \n", nextFeed.Url)

	err = s.db.MarkFeedFetched(s.ctx, nextFeed.ID)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	RSSFeed, err := fetchFeed(s.ctx, nextFeed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("%+v\n", RSSFeed)

	return nil

}
