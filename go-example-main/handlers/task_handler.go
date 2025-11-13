package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/acauhi/kanban-backend/models"
	"github.com/acauhi/kanban-backend/repository"
	"github.com/acauhi/kanban-backend/service"
)

const (
	pathTasks              = "/tasks/"
	msgInternalServerError = "Internal server error"
	msgInvalidRequestBody  = "Invalid request body"
	msgTaskIDRequired      = "Task ID is required"
	msgTaskNotFound        = "Task not found"
	msgMethodNotAllowed    = "Method not allowed"
)

type TaskHandler struct {
	service *service.TaskService
}

// NewTaskHandler cria uma nova instância do handler de tarefas
func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

// ServeHTTP roteia as requisições HTTP para os handlers apropriados
func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extrai o ID da tarefa da URL (se presente) sem causar panic para
	// caminhos como "/tasks" (sem barra final).
	// r.URL.Path pode ser "/tasks" ou "/tasks/123".
	p := strings.TrimPrefix(r.URL.Path, "/tasks")
	var id string
	if p == "" || p == "/" {
		id = ""
	} else {
		id = strings.TrimPrefix(p, "/")
	}

	switch r.Method {
	case http.MethodPost:
		if id == "" {
			h.handleCreate(w, r)
		} else {
			writeError(w, http.StatusMethodNotAllowed, msgMethodNotAllowed)
		}
	case http.MethodGet:
		if id == "" {
			h.handleGetAll(w, r)
		} else {
			h.handleGetByID(w, r)
		}
	case http.MethodPut:
		if id != "" {
			h.handleUpdate(w, r)
		} else {
			writeError(w, http.StatusMethodNotAllowed, msgMethodNotAllowed)
		}
	case http.MethodDelete:
		if id != "" {
			h.handleDelete(w, r)
		} else {
			writeError(w, http.StatusMethodNotAllowed, msgMethodNotAllowed)
		}
	default:
		writeError(w, http.StatusMethodNotAllowed, msgMethodNotAllowed)
	}
}

// handleCreate processa requisições POST para criar uma nova tarefa
func (h *TaskHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, msgInvalidRequestBody)
		return
	}

	task, err := h.service.CreateTask(req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidTitle) {
			writeError(w, http.StatusBadRequest, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, msgInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// handleGetAll processa requisições GET para listar todas as tarefas
func (h *TaskHandler) handleGetAll(w http.ResponseWriter, _ *http.Request) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		writeError(w, http.StatusInternalServerError, msgInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

// handleGetByID processa requisições GET para buscar uma tarefa por ID
func (h *TaskHandler) handleGetByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(pathTasks):]
	if id == "" {
		writeError(w, http.StatusBadRequest, msgTaskIDRequired)
		return
	}

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, msgTaskNotFound)
		} else {
			writeError(w, http.StatusInternalServerError, msgInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// handleUpdate processa requisições PUT para atualizar uma tarefa existente
func (h *TaskHandler) handleUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(pathTasks):]
	if id == "" {
		writeError(w, http.StatusBadRequest, msgTaskIDRequired)
		return
	}

	var req models.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, msgInvalidRequestBody)
		return
	}

	task, err := h.service.UpdateTask(id, req)
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, msgTaskNotFound)
		} else if errors.Is(err, service.ErrInvalidStatus) {
			writeError(w, http.StatusBadRequest, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, msgInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// handleDelete processa requisições DELETE para remover uma tarefa
func (h *TaskHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(pathTasks):]
	if id == "" {
		writeError(w, http.StatusBadRequest, msgTaskIDRequired)
		return
	}

	err := h.service.DeleteTask(id)
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, msgTaskNotFound)
		} else {
			writeError(w, http.StatusInternalServerError, msgInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// writeError escreve uma resposta de erro em JSON
func writeError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
