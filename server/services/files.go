package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Karakay-V/ShareLink/server/db"
)

type File struct {
	ID       int64  `json:"id"`
	FilePath string `json:"file_path"`
	Name     string `json:"name"`
	FileExt  string `json:"file_extension"`
}

// SaveFileService зберігає файл у файловій системі + пише в БД
func SaveFileService(lessonID int, fileName string, fileExt string, fileBytes []byte) (string, int64, error) {
	// генеруємо дату як папку
	dateFolder := time.Now().Format("02.01.2006") // приклад: 29.09.2025

	// створюємо шлях для папки
	root := "./uploads" // змінюй на "/" якщо хочеш абсолютний корінь
	lessonFolder := filepath.Join(root, dateFolder, fmt.Sprintf("%d", lessonID))

	// створюємо директорії якщо їх немає
	err := os.MkdirAll(lessonFolder, os.ModePerm)
	if err != nil {
		return "", 0, err
	}

	// 1️⃣ додаємо запис у БД щоб отримати id
	res, err := db.DB.Exec(`
        INSERT INTO files (file_path, name, file_extension)
        VALUES (?, ?, ?)`,
		"", fileName, fileExt,
	)
	if err != nil {
		return "", 0, err
	}

	fileID, err := res.LastInsertId()
	if err != nil {
		return "", 0, err
	}

	// 2️⃣ формуємо остаточний шлях
	filePath := filepath.Join(lessonFolder, fmt.Sprintf("%d.%s", fileID, fileExt))

	// 3️⃣ зберігаємо файл на диск
	err = os.WriteFile(filePath, fileBytes, 0644)
	if err != nil {
		return "", 0, err
	}

	// 4️⃣ оновлюємо запис у БД з реальним шляхом
	_, err = db.DB.Exec(`UPDATE files SET file_path = ? WHERE id = ?`, filePath, fileID)
	if err != nil {
		return "", 0, err
	}

	log.Printf("File saved: %s", filePath)
	return filePath, fileID, nil
}

// GetFileService повертає метадані файлу за id
func GetFileService(fileID int64) (*File, error) {
	row := db.DB.QueryRow(`SELECT id, file_path, name, file_extension FROM files WHERE id = ?`, fileID)

	var f File
	err := row.Scan(&f.ID, &f.FilePath, &f.Name, &f.FileExt)
	if err == sql.ErrNoRows {
		return nil, errors.New("file not found")
	} else if err != nil {
		return nil, err
	}

	return &f, nil
}
