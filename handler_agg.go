package main

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/lucianboboc/gator/internal/database"
	"github.com/lucianboboc/gator/internal/rss"
	"time"
)

const dateLayout = "Mon, 02 Jan 2006 15:04:05 -0700"

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("agg handlerAgg: time_between_reqs param required")
	}

	time_between_reqs := cmd.args[0]
	duration, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		err = scrapeFeeds(s.db)
		if err != nil {
			return err
		}
	}

	return nil
}

func scrapeFeeds(db *database.Queries) error {
	// get the next feed from DB
	dbFeed, err := db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	// mark it as fetched
	err = db.MarkFeedFetched(context.Background(), dbFeed.ID)
	if err != nil {
		return err
	}

	// fetch the feed with the url
	feed, err := rss.FetchFeed(context.Background(), dbFeed.Url)
	if err != nil {
		return err
	}

	// iterate over items in feed, print titles in the console
	for _, item := range feed.Channel.Items {
		publishedAt, err := time.Parse(dateLayout, item.PubDate)
		if err != nil {
			continue
		}

		now := time.Now()
		params := database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			FeedID:      dbFeed.ID,
			Description: item.Description,
			PublishedAt: publishedAt,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		_, err = db.CreatePost(context.Background(), params)
		if err != nil {
			continue
		}
	}

	return nil
}
