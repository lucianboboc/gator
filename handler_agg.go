package main

import (
	"context"
	"fmt"
	"github.com/lucianboboc/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := rss.FetchFeed(context.Background(), rss.DefaultFeedURL)
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
