package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/Karakay-V/ShareLink/server/db"
)

type Building struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetBuildings(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name FROM buildings")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var buildings []Building
	for rows.Next() {
		var b Building
		rows.Scan(&b.ID, &b.Name)
		buildings = append(buildings, b)
	}
	json.NewEncoder(w).Encode(buildings)
}
