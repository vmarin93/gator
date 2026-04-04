package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/vmarin93/gator/internal/database"
)

func handlerBrowsePosts(s *state, cmd command, user database.User) error {
	postsLimit := 2
	if len(cmd.args) == 1 {
		userLimit, err := strconv.Atoi(cmd.args[0])
		if err == nil {
			postsLimit = userLimit
		}
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(postsLimit),
	})
	if err != nil {
		return fmt.Errorf("Unable to retrieve posts from the db for user %s: %w", user.ID, err)
	}
	for _, post := range posts {
		fmt.Printf("%s\n", post.Title)
		fmt.Printf("%s\n", post.PublishedAt)
		fmt.Println("-------------------------------------")
		fmt.Printf("%s\n", post.Description)
		fmt.Println("=====================================")
	}
	return nil
}
