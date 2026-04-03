package main

import (
	"context"
	"fmt"
)

const feedURL = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("Unable to fetch the feed for the given url: %w", err)
	}
	printRSSFeed(feed)
	return nil
}

func printRSSFeed(feed *RSSFeed) {
	fmt.Printf(" * Title:      %v\n", feed.Channel.Title)
	fmt.Printf(" * Description:    %v\n", feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Printf("	- Title:	%v\n", item.Title)
		fmt.Printf("	- Content:	%v\n", item.Description)
		fmt.Printf("	- PubDate:	%v\n", item.PubDate)
	}
}
