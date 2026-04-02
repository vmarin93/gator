package main

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Expected username to be provided for the login command")
	}
	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		log.Fatal("Such a username doesn't exist in the db")
	}
	if err := s.conf.SetUser(user.Name); err != nil {
		return err
	}
	fmt.Printf("User %s has been logged in\n", user.Name)
	return nil
}
