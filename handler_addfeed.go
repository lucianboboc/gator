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

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("wrong number of arguments")
	}

	currentUser := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return err
	}

	feedName, feedURL := cmd.args[0], cmd.args[1]

	_, err = rss.FetchFeed(context.Background(), feedURL)
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
	return nil
}
