package models

type potroha struct {
	filename     string
	archive_size float64
	total_size   float64
	total_files  float64
	Files
}

type files struct {
	file_path string
	size      float64
	mimetype  string
}
