package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	currentUser := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return err
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}
	return nil
}
