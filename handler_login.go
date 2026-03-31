package main

import (
	"errors"
	"fmt"
)

func HandlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Expected username to be provided for the login command")
	}
	if err := s.conf.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Printf("User %s has been logged in\n", cmd.args[0])
	return nil
}
