package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vmarin93/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return errors.New("Please provide a name and a url for the feed you want to add")
	}
	dbFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("Unable to add feed to db: %w", err)
	}
	feedFollowData, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    dbFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("Failed to create FeedFollow whilst adding feed to feeds: %w", err)
	}
	fmt.Println("Feed created successfully:")
	printFeed(dbFeed, user)
	fmt.Println()
	fmt.Printf("Congratulations %s. You have just subscribed to %s\n", feedFollowData.UserName, feedFollowData.FeedName)
	fmt.Println("=====================================")
	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:		  %s\n", user.Name)
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to retrieve feeds from the db: %w", err)
	}
	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}
	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("Failed to retrieve user during listing feeds: %w", err)
		}
		printFeed(feed, user)
		fmt.Println("=====================================")
	}
	return nil
}
