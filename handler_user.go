package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/docherak/bd-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("login handler expects username argument, usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil && err == sql.ErrNoRows {
		log.Fatalf("user does not exist")
	}
	if err != nil {
		return err
	}

	err = s.config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("User has been set!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return errors.New("register handler expects username argument")
	}

	userID := uuid.New()
	username := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == nil {
		log.Fatal("user already exists")
	}

	_, err = s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        userID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      username,
	})
	if err != nil {
		return err
	}

	err = s.config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("User has been created!")

	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if s.config.CurrentUserName == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
