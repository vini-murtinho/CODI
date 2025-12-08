package repository

import (
	"testing"

	"github.com/acauhi/kanban-backend/models"
)

const (
	msgExpectedNoError         = "expected no error, got %v"
	msgExpectedErrTaskNotFound = "expected ErrTaskNotFound, got %v"
)

func TestInMemoryTaskRepositoryCreate(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	task := &models.Task{
		ID:     "1",
		Title:  "Test Task",
		Status: models.StatusTodo,
	}

	err := repo.Create(task)
	if err != nil {
		t.Fatalf(msgExpectedNoError, err)
	}

	retrieved, err := repo.GetByID("1")
	if err != nil {
		t.Fatalf(msgExpectedNoError, err)
	}

	if retrieved.Title != task.Title {
		t.Errorf("expected title %s, got %s", task.Title, retrieved.Title)
	}
}

func TestInMemoryTaskRepositoryGetByIDNotFound(t *testing.T) {
	repo := NewInMemoryTaskRepository()

	_, err := repo.GetByID("nonexistent")
	if err != ErrTaskNotFound {
		t.Errorf(msgExpectedErrTaskNotFound, err)
	}
}

func TestInMemoryTaskRepositoryUpdate(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	task := &models.Task{
		ID:     "1",
		Title:  "Original",
		Status: models.StatusTodo,
	}

	_ = repo.Create(task)

	task.Title = "Updated"
	err := repo.Update(task)
	if err != nil {
		t.Fatalf(msgExpectedNoError, err)
	}

	retrieved, _ := repo.GetByID("1")
	if retrieved.Title != "Updated" {
		t.Errorf("expected title Updated, got %s", retrieved.Title)
	}
}

func TestInMemoryTaskRepositoryDelete(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	task := &models.Task{
		ID:     "1",
		Title:  "Test",
		Status: models.StatusTodo,
	}

	_ = repo.Create(task)

	err := repo.Delete("1")
	if err != nil {
		t.Fatalf(msgExpectedNoError, err)
	}

	_, err = repo.GetByID("1")
	if err != ErrTaskNotFound {
		t.Errorf(msgExpectedErrTaskNotFound, err)
	}
}

func TestInMemoryTaskRepositoryGetAll(t *testing.T) {
	repo := NewInMemoryTaskRepository()

	tasks := []*models.Task{
		{ID: "1", Title: "Task 1", Status: models.StatusTodo},
		{ID: "2", Title: "Task 2", Status: models.StatusInProgress},
	}

	for _, task := range tasks {
		_ = repo.Create(task)
	}

	all, err := repo.GetAll()
	if err != nil {
		t.Fatalf(msgExpectedNoError, err)
	}

	if len(all) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(all))
	}
}

func TestInMemoryTaskRepositoryUpdateNotFound(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	task := &models.Task{
		ID:     "nonexistent",
		Title:  "Test",
		Status: models.StatusTodo,
	}

	err := repo.Update(task)
	if err != ErrTaskNotFound {
		t.Errorf(msgExpectedErrTaskNotFound, err)
	}
}

func TestInMemoryTaskRepositoryDeleteNotFound(t *testing.T) {
	repo := NewInMemoryTaskRepository()

	err := repo.Delete("nonexistent")
	if err != ErrTaskNotFound {
		t.Errorf(msgExpectedErrTaskNotFound, err)
	}
}
