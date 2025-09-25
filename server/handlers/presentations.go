package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"github.com/Karakay-V/ShareLink/server/db"
)

type Presentation struct {
	ID          int    `json:"id"`
	LessonID    int    `json:"lesson_id"`
	ClassroomID int    `json:"classroom_id"`
	Email       string `json:"email"`
	Description string `json:"description"`
	FilePath    string `json:"file_path"`
	UploadedAt  string `json:"uploaded_at"`
}

func HandlePresentations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPresentations(w, r)
	case http.MethodPost:
		createPresentation(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getPresentations(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(
		"SELECT id, lesson_id, classroom_id, email, description, file_path, uploaded_at FROM presentations")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var pres []Presentation
	for rows.Next() {
		var p Presentation
		rows.Scan(&p.ID, &p.LessonID, &p.ClassroomID, &p.Email, &p.Description, &p.FilePath, &p.UploadedAt)
		pres = append(pres, p)
	}
	json.NewEncoder(w).Encode(pres)
}

func createPresentation(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(20 << 20); err != nil { // 20MB
		http.Error(w, err.Error(), 400)
		return
	}

	lessonID := r.FormValue("lesson_id")
	classroomID := r.FormValue("classroom_id")
	email := r.FormValue("email")
	description := r.FormValue("description")

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file required", 400)
		return
	}
	defer file.Close()

	os.MkdirAll("uploads", os.ModePerm)
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)
	path := filepath.Join("uploads", filename)

	dst, _ := os.Create(path)
	defer dst.Close()
	io.Copy(dst, file)

	_, err = db.DB.Exec(
		"INSERT INTO presentations (lesson_id, classroom_id, email, description, file_path, uploaded_at) VALUES (?, ?, ?, ?, ?, ?)",
		lessonID, classroomID, email, description, path, time.Now().Format(time.RFC3339),
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Presentation uploaded successfully"))
}
