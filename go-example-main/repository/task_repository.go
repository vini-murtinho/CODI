package repository

import (
	"errors"
	"sync"

	"github.com/acauhi/kanban-backend/models"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskRepository interface {
	Create(task *models.Task) error
	GetAll() ([]*models.Task, error)
	GetByID(id string) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id string) error
}

type InMemoryTaskRepository struct {
	tasks map[string]*models.Task
	mu    sync.RWMutex
}

// NewInMemoryTaskRepository cria uma nova instância do repositório em memória
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]*models.Task),
		mu:    sync.RWMutex{},
	}
}

// Create adiciona uma nova tarefa ao repositório
func (r *InMemoryTaskRepository) Create(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[task.ID] = task
	return nil
}

// GetAll retorna todas as tarefas armazenadas
func (r *InMemoryTaskRepository) GetAll() ([]*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	tasks := make([]*models.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// GetByID busca uma tarefa específica pelo ID
func (r *InMemoryTaskRepository) GetByID(id string) (*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	task, exists := r.tasks[id]
	if !exists {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

// Update atualiza uma tarefa existente no repositório
func (r *InMemoryTaskRepository) Update(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.tasks[task.ID]; !exists {
		return ErrTaskNotFound
	}
	r.tasks[task.ID] = task
	return nil
}

// Delete remove uma tarefa do repositório pelo ID
func (r *InMemoryTaskRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.tasks[id]; !exists {
		return ErrTaskNotFound
	}
	delete(r.tasks, id)
	return nil
}
