package usecase

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"mime"
	"path/filepath"
	"task-doodocs/models"
)

type ArchiveUsecase struct{}

func NewArchiveUseCase() *ArchiveUsecase {
	return &ArchiveUsecase{}
}

func (uc *ArchiveUsecase) ProcessArchive(file []byte, filename string) (*models.ArchiveInfo, error) {
	reader, err := zip.NewReader(bytes.NewReader(file), int64(len(file)))
	if err != nil {
		//log suda nujen
		return nil, errors.New("Provided file is not a valid ZIP archive")
	}

	var totalSize float64
	var files []models.FileInfo

	for _, f := range reader.File {
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}

		var fileSize float64
		if stat, err := f.FileInfo().Sys().(*zip.FileHeader); err == false {
			fileSize = float64(stat.UncompressedSize64)
		} else {
			buf := new(bytes.Buffer)
			_, _ = io.Copy(buf, rc)
			fileSize = float64(buf.Len())
		}
		rc.Close()

		files = append(files, models.FileInfo{
			FilePath: f.Name,
			Size:     fileSize,
			MimeType: mime.TypeByExtension(filepath.Ext(f.Name)),
		})
		totalSize += fileSize
	}
	return &models.ArchiveInfo{
		Filename:    filename,
		ArchiveSize: float64(len(file)),
		TotalSize:   totalSize,
		TotalFiles:  len(file),
		Files:       files,
	}, nil
}
