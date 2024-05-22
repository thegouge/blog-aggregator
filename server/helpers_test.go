package main

import (
	"testing"
)

func TestFetchingFeedData(t *testing.T) {
	feedResult, err := fetchFeedData("https://blog.boot.dev/index.xml")

	if err != nil {
		t.Fatalf("Fetching errored out wiht error %v\n", err)
	}

	var testValues = []struct {
		name  string
		input string
		want  string
	}{
		{"Testing the Title", feedResult.Channel.Title, "Boot.dev Blog"},
		{"Testing the Description", feedResult.Channel.Description, "Recent content on Boot.dev Blog"},
		{"Testing a random item", feedResult.Channel.Item[22].Title, "The One Thing I'd Change About Go"},
	}
	for _, tt := range testValues {
		println(tt.name)
		if tt.input != tt.want {
			t.Errorf("got %s, wanted %s", tt.input, tt.want)
		}
	}
}
