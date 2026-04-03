package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Expected username to be provided for the login command")
	}
	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Could not get user %s from db: %w", cmd.args[0], err)
	}
	if err := s.conf.SetUser(user.Name); err != nil {
		return fmt.Errorf("Could set user in conf during login %w", err)
	}
	fmt.Printf("User %s has been logged in\n", user.Name)
	return nil
}
