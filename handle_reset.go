package main

import (
	"context"
	"log"
)

func handleReset(s *state, cmd command) error {
	if err := s.db.ResetUsers(context.Background()); err != nil {
		log.Fatalf("Unable to reset the users table in the db: %v", err)
	}
	log.Print("Users table in the db was reset succesfully!")
	return nil
}
