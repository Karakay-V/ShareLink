package main

import (
	"log"
	"net/http"

	"github.com/Karakay-V/ShareLink/server/db"
	"github.com/Karakay-V/ShareLink/server/handlers"
)

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	db.InitDB()
	defer db.DB.Close()

	mux := http.NewServeMux()

	// твої ендпоінти
	mux.HandleFunc("/buildings", handlers.GetBuildings)
	mux.HandleFunc("/classrooms", handlers.GetClassrooms)
	mux.HandleFunc("/lessons", handlers.GetLessons)
	mux.HandleFunc("/presentations", handlers.HandlePresentations)

	// обгортаємо middleware
	handler := corsMiddleware(mux)

	log.Println("Server started at :8002")
	log.Fatal(http.ListenAndServe(":8002", handler))
}
