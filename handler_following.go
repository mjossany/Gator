package main

import (
	"context"
	"fmt"

	"github.com/mjossany/Gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows for user: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("You are not following any feeds.")
		return nil
	}

	fmt.Printf("You are following %d feed(s):\n", len(feedFollows))
	for _, feedFollow := range feedFollows {
		fmt.Printf("- %s\n", feedFollow.FeedName)
	}

	return nil
}
