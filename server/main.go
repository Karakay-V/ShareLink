package main

import (
	"log"
	"net/http"

	"github.com/Karakay-V/ShareLink/server/config"
	"github.com/Karakay-V/ShareLink/server/db"
	"github.com/Karakay-V/ShareLink/server/handlers"
	"github.com/Karakay-V/ShareLink/server/middleware"
	"github.com/Karakay-V/ShareLink/server/services"
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
	// init configs
	config.InitConfig()

	db.InitDB()
	defer db.DB.Close()

	services.StartOutdateMarkerService()

	mux := http.NewServeMux()

	// твої ендпоінти
	mux.HandleFunc("/buildings", handlers.GetBuildings)
	mux.HandleFunc("/classrooms", handlers.GetClassrooms)
	mux.HandleFunc("/lessons", handlers.GetLessons)

	mux.HandleFunc("/login", handlers.TeacherLogin)
	mux.Handle(
		"/register",
		middleware.AuthMiddleware(http.HandlerFunc(handlers.RegisterTeacher)))

	mux.Handle(
		"/presentations", 
		http.RedirectHandler("/presentations/", http.StatusMovedPermanently))
	mux.Handle("/presentations/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// POST без авторизації
			handlers.HandlePresentations(w, r)
		} else {
			// решта методів — з авторизацією
			middleware.AuthMiddleware(http.HandlerFunc(handlers.HandlePresentations)).ServeHTTP(w, r)
		}
	}))

	// обгортаємо middleware
	handler := corsMiddleware(mux)

	log.Println("Server started at :8002")
	log.Fatal(http.ListenAndServe(":8002", handler))
}
