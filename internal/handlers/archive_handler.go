package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"task-doodocs/models"
)

type ArchiveHandler struct {
	usecase models.ArchiveUsecase
}

func NewArchiveHandler(usecase models.ArchiveUsecase) *ArchiveHandler {
	return &ArchiveHandler{usecase: usecase}
}

func (h *ArchiveHandler) ProcessArchive(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "failed to retrieve file from request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to read file", http.StatusInternalServerError)
		return
	}

	archiveInfo, err := h.usecase.ProcessArchive(fileBytes, header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(archiveInfo)
}
