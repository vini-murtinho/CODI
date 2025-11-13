package main

import (
	"log"
	"net/http"

	"github.com/acauhi/kanban-backend/handlers"
	"github.com/acauhi/kanban-backend/repository"
	"github.com/acauhi/kanban-backend/service"
)

// main inicializa o servidor HTTP com todas as dependências
func main() {
	repo := repository.NewInMemoryTaskRepository()
	svc := service.NewTaskService(repo)
	handler := handlers.NewTaskHandler(svc)

	mux := http.NewServeMux()
	mux.Handle("/tasks", corsMiddleware(handler))
	mux.Handle("/tasks/", corsMiddleware(handler))

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

// corsMiddleware adiciona headers CORS para permitir requisições do frontend
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allowing all origins in development for convenience. Adjust for production.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
