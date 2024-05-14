package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

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

	_, err := db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:            feed.ID,
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		log.Println("Error marking feed as fetched:", err)
		return
	}

	data, err := fetchFeedData(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed %s, %v", feed.Name, err)
	}

	for _, item := range data.Channel.Item {
		log.Println("New item! -", item.Title, "on feed", feed.Name)
	}
}
