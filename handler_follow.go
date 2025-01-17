package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lucianboboc/gator/internal/database"
	"time"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("following url is required")
	}

	url := cmd.args[0]
	dbFeed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("follow GetFeedByUrl: %w", err)
	}

	currentUser := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("follow GetUser: %w", err)
	}

	now := time.Now()
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    dbFeed.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("follow CreateFeedFollow: %w", err)
	}

	fmt.Println(dbFeed.Name)
	fmt.Println(currentUser)

	return nil
}
