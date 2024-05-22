package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s\n", err)
		w.WriteHeader(500)
		return
	}
	if status != 200 {
		w.WriteHeader(status)
	}

	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func fetchFeedData(url string) (RssFeed, error) {
	response, err := http.Get(url)

	if err != nil {
		return RssFeed{}, err
	}

	defer response.Body.Close()
	if response.StatusCode > 399 {
		return RssFeed{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return RssFeed{}, err
	}

	responseData := RssFeed{}

	err = xml.Unmarshal(body, &responseData)

	if err != nil {
		return RssFeed{}, err
	}

	return responseData, nil
}
