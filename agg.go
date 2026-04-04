package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	httpClient := http.DefaultClient
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agent", "gator")
	res, err := httpClient.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer res.Body.Close()
	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	feed := RSSFeed{}
	if err := xml.Unmarshal(resData, &feed); err != nil {
		return &RSSFeed{}, err
	}
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feed.Channel.Item[i] = item
	}
	return &feed, nil
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to retrieve the next feed to fetch from db: %w", err)
	}
	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("Failed to fetch feed %s: %w", feed.Name, err)
	}
	if err := s.db.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		return fmt.Errorf("Unable to mark feed %s as fetched in the db: %w", feed.Name, err)
	}
	printRSSFeed(rssFeed)
	return nil
}
