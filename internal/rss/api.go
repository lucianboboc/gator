package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

const (
	DefaultFeedURL = "https://www.wagslane.dev/index.xml"
)

func FetchFeed(ctx context.Context, feedURL string) (*Feed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var feed Feed
	err = xml.Unmarshal(bodyBytes, &feed)
	if err != nil {
		return nil, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := 0; i < len(feed.Channel.Items); i++ {
		item := feed.Channel.Items[i]
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feed.Channel.Items[i] = item
	}

	return &feed, nil
}
