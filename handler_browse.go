package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lucianboboc/gator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		limitArg, err := strconv.ParseInt(cmd.args[0], 10, 64)
		if err == nil {
			limit = int(limitArg)
		}
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return errors.New("handlerBrowse: can't fetch posts for user")
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Description)
	}

	return nil
}
