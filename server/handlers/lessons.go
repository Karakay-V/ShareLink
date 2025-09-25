package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Karakay-V/ShareLink/server/db"
)

type Lesson struct {
	ID         int    `json:"id"`
	PairNumber int    `json:"pair_number"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

func GetLessons(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, pair_number, start_time, end_time FROM lessons")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var lessons []Lesson
	for rows.Next() {
		var l Lesson
		rows.Scan(&l.ID, &l.PairNumber, &l.StartTime, &l.EndTime)
		lessons = append(lessons, l)
	}
	json.NewEncoder(w).Encode(lessons)
}
