package services

import (
	"log"
	"time"

	"github.com/Karakay-V/ShareLink/server/db"
)

func StartOutdateMarkerService() {
	ticker := time.NewTicker(1 * time.Minute) // перевірка кожну хвилину
	go func() {
		for range ticker.C {
			markOutdated()
		}
	}()
}

func markOutdated() {
	now := time.Now().Format("15:04")

	query := `
        UPDATE presentations
        SET is_outdated = 1
        WHERE lesson_id IN (
            SELECT id FROM lessons
            WHERE end_time <= ?
        ) AND is_outdated = 0
    `

	res, err := db.DB.Exec(query, now)
	if err != nil {
		log.Println("Error marking outdated presentations:", err)
		return
	}

	affected, _ := res.RowsAffected()
	if affected > 0 {
		log.Println("Marked presentations as outdated:", affected, "at", now)
	}
}
