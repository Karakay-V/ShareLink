package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Karakay-V/ShareLink/server/db"
	"github.com/Karakay-V/ShareLink/server/services"
)

type Presentation struct {
	ID            int    `json:"id"`
	LessonID      int    `json:"lesson_id,omitempty"`      // внутрішній
	PairNumber    string `json:"pair_number,omitempty"`    // для фронту
	ClassroomID   int    `json:"classroom_id,omitempty"`   // внутрішній
	ClassroomName string `json:"classroom_name,omitempty"` // для фронту
	Email         string `json:"email"`
	Description   string `json:"description"`
	FileID        int    `json:"file_id"`   // 🔑 id файлу
	FileName      string `json:"file_name"` // 🔑 замість file_path
	UploadedAt    string `json:"uploaded_at"`
	IsOutdated    bool   `json:"is_outdated"`
}

func HandlePresentations(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch {
	// /presentations/all
	case strings.HasPrefix(path, "/presentations/all"):
		getAllPresentations(w, r)

	// /presentations/get/lesson/{lesson_number}
	case strings.HasPrefix(path, "/presentations/get/lesson/"):
		getByLesson(w, r)

	// /presentations/download/{file_id}
	case strings.HasPrefix(path, "/presentations/download/"):
		downloadFile(w, r)

	// POST /presentations
	case r.Method == http.MethodPost:
		createPresentation(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ===================== GET ALL =====================
func getAllPresentations(w http.ResponseWriter, r *http.Request) {
	showHidden := r.URL.Query().Get("show_hidden") == "true"

	query := `
		SELECT p.id, l.pair_number, c.name, p.email, p.description, p.file_id, f.name, p.uploaded_at, p.is_outdated
		FROM presentations p
		JOIN lessons l ON p.lesson_id = l.id
		JOIN classrooms c ON p.classroom_id = c.id
		LEFT JOIN files f ON p.file_id = f.id
	`
	if !showHidden {
		query += " WHERE is_outdated = 0"
	}

	rows, err := db.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var pres []Presentation
	for rows.Next() {
		var p Presentation
		rows.Scan(&p.ID, &p.PairNumber, &p.ClassroomName, &p.Email, &p.Description, &p.FileID, &p.FileName, &p.UploadedAt, &p.IsOutdated)
		pres = append(pres, p)
	}
	json.NewEncoder(w).Encode(pres)
}

// ===================== GET BY LESSON =====================
func getByLesson(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		http.Error(w, "pair_number required", 400)
		return
	}
	pairNumber := parts[4]
	classroomName := r.URL.Query().Get("classroom")
	if classroomName == "" {
		http.Error(w, "classroom name required", 400)
		return
	}

	query := `
		SELECT p.id, l.pair_number, c.name, p.email, p.description, p.file_id, f.name, p.uploaded_at, p.is_outdated
		FROM presentations p
		JOIN lessons l ON p.lesson_id = l.id
		JOIN classrooms c ON p.classroom_id = c.id
		LEFT JOIN files f ON p.file_id = f.id
		WHERE l.pair_number = ? AND c.name = ?
	`

	rows, err := db.DB.Query(query, pairNumber, classroomName)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var pres []Presentation
	for rows.Next() {
		var p Presentation
		rows.Scan(&p.ID, &p.PairNumber, &p.ClassroomName, &p.Email, &p.Description, &p.FileID, &p.FileName, &p.UploadedAt, &p.IsOutdated)
		if !p.IsOutdated {
			pres = append(pres, p)
		}
	}
	json.NewEncoder(w).Encode(pres)
}

// ===================== DOWNLOAD FILE =====================
func downloadFile(w http.ResponseWriter, r *http.Request) {
	// /presentations/download/{file_id}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "file_id required", 400)
		return
	}
	fileID := parts[3]

	var filePath, fileName string
	err := db.DB.QueryRow(`SELECT file_path, name FROM files WHERE id = ?`, fileID).
		Scan(&filePath, &fileName)
	if err == sql.ErrNoRows {
		http.Error(w, "file not found", 404)
		return
	} else if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "cannot open file", 500)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")
	io.Copy(w, file)
}

// ===================== CREATE =====================
func createPresentation(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(20 << 20); err != nil { // 20MB
		http.Error(w, err.Error(), 400)
		return
	}

	lessonIDStr := r.FormValue("lesson_id")
	classroomIDStr := r.FormValue("classroom_id")
	email := r.FormValue("email")
	description := r.FormValue("description")

	lessonID, err := strconv.Atoi(lessonIDStr)
	if err != nil {
		http.Error(w, "invalid lesson_id", 400)
		return
	}
	classroomID, err := strconv.Atoi(classroomIDStr)
	if err != nil {
		http.Error(w, "invalid classroom_id", 400)
		return
	}

	// отримуємо файл
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file required", 400)
		return
	}
	defer file.Close()

	// читаємо байти файлу
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "cannot read file", 500)
		return
	}

	// визначаємо розширення
	ext := filepath.Ext(handler.Filename)
	if len(ext) > 0 {
		ext = ext[1:] // забираємо "."
	}

	// зберігаємо файл через сервіс
	_, fileID, err := services.SaveFileService(lessonID, handler.Filename, ext, fileBytes)
	if err != nil {
		http.Error(w, "cannot save file", 500)
		return
	}

	// тепер створюємо запис у presentations
	_, err = db.DB.Exec(
		`INSERT INTO presentations 
		 (lesson_id, classroom_id, email, description, file_id, uploaded_at, is_outdated) 
		 VALUES (?, ?, ?, ?, ?, ?, 0)`,
		lessonID, classroomID, email, description, fileID, time.Now().Format(time.RFC3339),
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Presentation uploaded successfully with file_id=%d", fileID)))
}
