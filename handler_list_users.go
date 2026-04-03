package main

import (
	"context"
	"fmt"
)

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to retrieve list of users from db: %w", err)
	}
	for i := range len(users) {
		if users[i].Name == s.conf.CurrentUserName {
			fmt.Printf("* %s (current)\n", users[i].Name)
		} else {
			fmt.Printf("* %s\n", users[i].Name)
		}
	}
	return nil
}
