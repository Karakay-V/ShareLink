package models

type File struct {
	ID            int    `json:"id"`
	FilePath      string `json:"file_path"`
	Name          string `json:"name"`
	FileExtension string `json:"file_extension"`
}
