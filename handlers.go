package main

import "net/http"

func healthHandler(w http.ResponseWriter, Request *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func errHandler(w http.ResponseWriter, Request *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
