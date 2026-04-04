package main

import (
	"context"
	"fmt"

	"github.com/vmarin93/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Failure to authenticate user in middlewareLoggedIn operation, %w",
				err)
		}
		return handler(s, cmd, user)
	}
}
