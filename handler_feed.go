package main

import (
	"context"
	"fmt"
	"time"

	"github.com/docherak/bd-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("addfeed handler expects feed name and feed url arguments, usage: %s <feed_name> <feed_url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}

	feedID := uuid.New()
	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]
	userID := user.ID

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        feedID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    userID,
	})
	if err != nil {
		return err
	}

	ffID := uuid.New()

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        ffID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println("---")
		fmt.Printf("feed: %s\n", feed.Feed.Name)
		fmt.Printf("url: %s\n", feed.Feed.Url)
		fmt.Printf("created by: %s\n", feed.User.Name)
	}
	fmt.Println("---")

	return nil
}
