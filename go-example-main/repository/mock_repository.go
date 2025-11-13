package repository

import (
	"errors"

	"github.com/acauhi/kanban-backend/models"
)

type MockTaskRepository struct {
	CreateFunc  func(task *models.Task) error
	GetAllFunc  func() ([]*models.Task, error)
	GetByIDFunc func(id string) (*models.Task, error)
	UpdateFunc  func(task *models.Task) error
	DeleteFunc  func(id string) error
}

// Create executa a função mock de criação se definida
func (m *MockTaskRepository) Create(task *models.Task) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(task)
	}
	return nil
}

// GetAll executa a função mock de listagem se definida
func (m *MockTaskRepository) GetAll() ([]*models.Task, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return nil, nil
}

// GetByID executa a função mock de busca por ID se definida
func (m *MockTaskRepository) GetByID(id string) (*models.Task, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return nil, nil
}

// Update executa a função mock de atualização se definida
func (m *MockTaskRepository) Update(task *models.Task) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(task)
	}
	return nil
}

// Delete executa a função mock de remoção se definida
func (m *MockTaskRepository) Delete(id string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}

var ErrMockError = errors.New("mock error")
