package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAggregate(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	requestsDelta, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	ticker := time.NewTicker(requestsDelta)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {

	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	feed, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		println(item.Title)
	}

	return nil
}
