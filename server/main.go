package main

import (
	"log"
	"net/http"

	"github.com/Karakay-V/ShareLink/server/db"
	"github.com/Karakay-V/ShareLink/server/handlers"
)

func main() {
	db.InitDB()
	defer db.DB.Close()

	http.HandleFunc("/buildings", handlers.GetBuildings)
	http.HandleFunc("/classrooms", handlers.GetClassrooms)
	http.HandleFunc("/lessons", handlers.GetLessons)
	http.HandleFunc("/presentations", handlers.HandlePresentations)

	log.Println("Server started at :8002")
	log.Fatal(http.ListenAndServe(":8002", nil))
}
