package main

import (
	"context"
	"errors"
	"github.com/lucianboboc/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("unfollow handlerUnfollow: feed url required")
	}

	params := database.DeleteFeedByUrlForUserParams{
		UserID: user.ID,
		Url:    cmd.args[0],
	}
	return s.db.DeleteFeedByUrlForUser(context.Background(), params)
}
