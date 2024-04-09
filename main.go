package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/thegouge/blog-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT environment variable is not set!")
	}
	dbURL := os.Getenv("DB_CONNECTOR")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error trying to load database")
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{DB: dbQueries}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/users", apiCfg.HandleUserCreate)

	mux.HandleFunc("GET /v1/healthz", healthHandler)
	mux.HandleFunc("GET /v1/err", errHandler)

	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: corsMux,
	}

	log.Printf("Serving on port : %s\n", PORT)
	log.Fatal(srv.ListenAndServe())
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
