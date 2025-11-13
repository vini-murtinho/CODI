package service

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/acauhi/kanban-backend/models"
	"github.com/acauhi/kanban-backend/repository"
)

var (
	ErrInvalidTitle  = errors.New("title is required")
	ErrInvalidStatus = errors.New("invalid status")
)

// idCounter helps ensure unique IDs when created in rapid succession during tests
var idCounter int64

type TaskService struct {
	repo repository.TaskRepository
}

// NewTaskService cria uma nova instância do serviço de tarefas
func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

// CreateTask cria uma nova tarefa com status inicial "todo"
func (s *TaskService) CreateTask(req models.CreateTaskRequest) (*models.Task, error) {
	if req.Title == "" {
		return nil, ErrInvalidTitle
	}

	// Gera ID único usando timestamp + UUID
	id := generateID()

	task := &models.Task{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Status:      models.StatusTodo,
		Completed:   false,
	}

	if err := s.repo.Create(task); err != nil {
		return nil, err
	}

	return task, nil
}

// GetAllTasks retorna todas as tarefas cadastradas
func (s *TaskService) GetAllTasks() ([]*models.Task, error) {
	return s.repo.GetAll()
}

// GetTaskByID busca uma tarefa específica pelo ID
func (s *TaskService) GetTaskByID(id string) (*models.Task, error) {
	return s.repo.GetByID(id)
}

// UpdateTask atualiza os campos de uma tarefa existente
func (s *TaskService) UpdateTask(id string, req models.UpdateTaskRequest) (*models.Task, error) {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		if *req.Title == "" {
			return nil, ErrInvalidTitle
		}
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Status != nil {
		if !isValidStatus(*req.Status) {
			return nil, ErrInvalidStatus
		}
		task.Status = *req.Status
		// Atualiza o campo 'completed' baseado no status
		task.Completed = (*req.Status == models.StatusDone)
	}

	if err := s.repo.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}

// DeleteTask remove uma tarefa pelo ID
func (s *TaskService) DeleteTask(id string) error {
	return s.repo.Delete(id)
}

// isValidStatus valida se o status fornecido é um dos valores permitidos
func isValidStatus(status models.Status) bool {
	switch status {
	case models.StatusTodo, models.StatusInProgress, models.StatusDone:
		return true
	default:
		return false
	}
}

// generateID cria um ID único para uma tarefa
func generateID() string {
	// combine timestamp with an atomic counter to avoid collisions in tests
	cnt := atomic.AddInt64(&idCounter, 1)
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), cnt)
}
