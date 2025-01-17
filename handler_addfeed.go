package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lucianboboc/gator/internal/database"
	"github.com/lucianboboc/gator/internal/rss"
	"time"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return errors.New("wrong number of arguments")
	}

	feedName, feedURL := cmd.args[0], cmd.args[1]

	_, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	now := time.Now()
	createFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	dbFeed, err := s.db.CreateFeed(context.Background(), createFeed)
	if err != nil {
		return err
	}
	fmt.Println(dbFeed)

	params := database.CreateFeedFollowParams{
		ID:        dbFeed.ID,
		UserID:    user.ID,
		FeedID:    dbFeed.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	return nil
}
