package main

import (
	"context"
	"fmt"
	"log"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		log.Fatal("failed to delete users")
	}
	fmt.Println("Users deleted!")

	return nil
}
