package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Karakay-V/ShareLink/server/db"
)

type Classroom struct {
	ID         int    `json:"id"`
	BuildingID int    `json:"building_id"`
	Name       string `json:"name"`
}

func GetClassrooms(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, building_id, name FROM classrooms")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var classrooms []Classroom
	for rows.Next() {
		var c Classroom
		rows.Scan(&c.ID, &c.BuildingID, &c.Name)
		classrooms = append(classrooms, c)
	}
	json.NewEncoder(w).Encode(classrooms)
}
