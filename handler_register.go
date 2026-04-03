package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/vmarin93/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Expected username to be provided for the register command")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0]},
	)
	if err != nil {
		pqErr, ok := errors.AsType[*pq.Error](err)
		if ok {
			if pqErr.Code == "23505" {
				log.Fatal("A user with that name already exists in the db")
			}
		}
		return err
	}
	println("User was created succesfully!")
	printUser(user)
	if err := s.conf.SetUser(user.Name); err != nil {
		return fmt.Errorf("Could not set user in conf during registration: %w", err)
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
