package main

import (
	"fmt"
	"log"
	"net/http"
	"task-doodocs/internal/handlers"
	"task-doodocs/internal/usecase"
)

func main() {
	ArchiveUsecase := usecase.NewArchiveUseCase()
	ArchiveHandler := handlers.NewArchiveHandler(ArchiveUsecase)

	http.HandleFunc("/api/archive/information", ArchiveHandler.ProcessArchive)

	port := 8080
	fmt.Printf("Server is running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
