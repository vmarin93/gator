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

func handleReset(s *state, cmd command) error {
	if err := s.db.ResetUsers(context.Background()); err != nil {
		log.Fatalf("Unable to reset the users table in the db: %v", err)
	}
	log.Print("Users table in the db was reset succesfully!")
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
