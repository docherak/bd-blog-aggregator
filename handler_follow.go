package main

import (
	"context"
	"fmt"
	"time"

	"github.com/docherak/bd-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerListFollows(s *state, cmd command, user database.User) error {

	userID := user.ID

	ffs, err := s.db.ListFeedFollow(context.Background(), userID)
	if err != nil {
		return err
	}

	fmt.Printf("Feeds followed by user %s\n", user.Name)
	for _, ff := range ffs {
		fmt.Printf("- %s\n", ff.FeedName)
	}

	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("follow handler expects a feed url argument, usage: %s <feed_url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	ffID := uuid.New()

	ff, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        ffID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Follow created for user %s and feed %s\n", ff.UserName, ff.FeedName)

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("unfollow handler expects a feed url argument, usage: %s <feed_url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("User %s unfollowd feed %s\n", user.Name, feed.Name)

	return nil
}
