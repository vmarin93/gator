package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vmarin93/gator/internal/database"
)

func handlerFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("Please provide a url for the feed you want to follow")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Unable to retrieve feed from db whilst subscribing to feed : %w", err)
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Unable to subscribe to feed: %w", err)
	}
	fmt.Printf("Congratulations %s. You have just subscribed to %s",
		feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handlerListFeedFollows(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.ListFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Unable to retrieve the list of feeds the user follows: %w", err)
	}
	for _, feedFollow := range feedFollows {
		fmt.Printf("* %s\n", feedFollow.FeedName)
	}
	return nil
}
