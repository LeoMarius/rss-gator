package main

import (
	"fmt"
)

func handlerAgg(s *state, cmd command) error {

	url := "https://www.wagslane.dev/index.xml"

	RSSFeed, err := fetchFeed(s.ctx, url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("%+v\n", RSSFeed)

	return nil
}
