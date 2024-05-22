package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/thegouge/blog-aggregator/internal/database"
)

func startScraping(db *database.Queries, numThreads int, timeBetweenRequests time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", numThreads, timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(numThreads))
		if err != nil {
			log.Println("error fetching feeds:", err)
			continue
		}

		waitGroup := &sync.WaitGroup{}

		for _, feed := range feeds {
			waitGroup.Add(1)
			go scrapeFeed(waitGroup, feed, db)
		}
		waitGroup.Wait()
	}
}

func scrapeFeed(waitGroup *sync.WaitGroup, feed database.Feed, db *database.Queries) {
	defer waitGroup.Done()

	sqlNow := sql.NullTime{Time: time.Now(), Valid: true}
	_, err := db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:            feed.ID,
		LastFetchedAt: sqlNow,
	})
	if err != nil {
		log.Println("Error marking feed as fetched:", err)
		return
	}

	data, err := fetchFeedData(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed %s, %v", feed.Name, err)
	}

	newID := uuid.NewString()
	for _, item := range data.Channel.Item {
		publishedTime, err := time.Parse(time.RFC1123Z, item.PubDate)

		if err != nil {
			log.Printf("Couldn't parses date %v, with err %v", item.PubDate, err)
		}

		description := sql.NullString{String: "", Valid: false}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          newID,
			CreatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: publishedTime,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("Failed to create post:", err)
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(data.Channel.Item))
}
