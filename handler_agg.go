package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const feedURL = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("Please provide with an interval at which to aggregate feeds")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Unable to parse argument to time.Duration object: %w", err)
	}
	ticker := time.NewTicker(timeBetweenRequests)
	log.Printf("Collecting feeds every %s...", timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func printRSSFeed(feed *RSSFeed) {
	// fmt.Printf(" * Title:      %v\n", feed.Channel.Title)
	// fmt.Printf(" * Description:    %v\n", feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Printf("	- Title:	%v\n", item.Title)
		// fmt.Printf("	- Content:	%v\n", item.Description)
		// fmt.Printf("	- PubDate:	%v\n", item.PubDate)
	}
}
