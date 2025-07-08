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

	_, err = s.db.CreateFeed(context.Background(), database.CreateFeedParams{
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

	//fmt.Printf("Feed created: +%v\n", feed)

	return nil
}
