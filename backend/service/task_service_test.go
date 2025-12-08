package service

import (
	"testing"

	"github.com/acauhi/kanban-backend/models"
	"github.com/acauhi/kanban-backend/repository"
)

const msgExpectedNoError = "expected no error, got %v"

func TestTaskServiceCreateTask(t *testing.T) {
	repo := repository.NewInMemoryTaskRepository()
	svc := NewTaskService(repo)

	req := models.CreateTaskRequest{
		Title:       "New Task",
		Description: "Description",
	}

	task, err := svc.CreateTask(req)
	if err != nil {
		t.Fatalf(msgExpectedNoError, err)
	}

	if task.Title != req.Title {
		t.Errorf("expected title %s, got %s", req.Title, task.Title)
	}

	if task.Status != models.StatusTodo {
		t.Errorf("expected status %s, got %s", models.StatusTodo, task.Status)
	}
}

func TestTaskServiceCreateTaskEmptyTitle(t *testing.T) {
	repo := repository.NewInMemoryTaskRepository()
	svc := NewTaskService(repo)

	req := models.CreateTaskRequest{
		Title: "",
	}

	_, err := svc.CreateTask(req)
	if err != ErrInvalidTitle {
		t.Errorf("expected ErrInvalidTitle, got %v", err)
	}
}

func TestTaskServiceUpdateTask(t *testing.T) {
	repo := repository.NewInMemoryTaskRepository()
	svc := NewTaskService(repo)

	task, _ := svc.CreateTask(models.CreateTaskRequest{Title: "Original"})

	newTitle := "Updated"
	newStatus := models.StatusInProgress
	req := models.UpdateTaskRequest{
		Title:  &newTitle,
		Status: &newStatus,
	}

	updated, err := svc.UpdateTask(task.ID, req)
	if err != nil {
		t.Fatalf(msgExpectedNoError, err)
	}

	if updated.Title != newTitle {
		t.Errorf("expected title %s, got %s", newTitle, updated.Title)
	}

	if updated.Status != newStatus {
		t.Errorf("expected status %s, got %s", newStatus, updated.Status)
	}
}

func TestTaskServiceUpdateTaskInvalidStatus(t *testing.T) {
	repo := repository.NewInMemoryTaskRepository()
	svc := NewTaskService(repo)

	task, _ := svc.CreateTask(models.CreateTaskRequest{Title: "Test"})

	invalidStatus := models.Status("invalid")
	req := models.UpdateTaskRequest{
		Status: &invalidStatus,
	}

	_, err := svc.UpdateTask(task.ID, req)
	if err != ErrInvalidStatus {
		t.Errorf("expected ErrInvalidStatus, got %v", err)
	}
}

func TestTaskServiceUpdateTaskNotFound(t *testing.T) {
	repo := repository.NewInMemoryTaskRepository()
	svc := NewTaskService(repo)

	newTitle := "Updated"
	req := models.UpdateTaskRequest{Title: &newTitle}

	_, err := svc.UpdateTask("nonexistent", req)
	if err != repository.ErrTaskNotFound {
		t.Errorf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestTaskServiceDeleteTask(t *testing.T) {
	repo := repository.NewInMemoryTaskRepository()
	svc := NewTaskService(repo)

	task, _ := svc.CreateTask(models.CreateTaskRequest{Title: "Test"})

	err := svc.DeleteTask(task.ID)
	if err != nil {
		t.Fatalf(msgExpectedNoError, err)
	}

	_, err = svc.GetTaskByID(task.ID)
	if err != repository.ErrTaskNotFound {
		t.Errorf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestTaskServiceGetAllTasks(t *testing.T) {
	repo := repository.NewInMemoryTaskRepository()
	svc := NewTaskService(repo)

	_, _ = svc.CreateTask(models.CreateTaskRequest{Title: "Task 1"})
	_, _ = svc.CreateTask(models.CreateTaskRequest{Title: "Task 2"})

	tasks, err := svc.GetAllTasks()
	if err != nil {
		t.Fatalf(msgExpectedNoError, err)
	}

	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(tasks))
	}
}

func TestTaskServiceUpdateTaskEmptyTitle(t *testing.T) {
	repo := repository.NewInMemoryTaskRepository()
	svc := NewTaskService(repo)

	task, _ := svc.CreateTask(models.CreateTaskRequest{Title: "Test"})

	emptyTitle := ""
	req := models.UpdateTaskRequest{Title: &emptyTitle}

	_, err := svc.UpdateTask(task.ID, req)
	if err != ErrInvalidTitle {
		t.Errorf("expected ErrInvalidTitle, got %v", err)
	}
}

func TestTaskServiceUpdateTaskDescription(t *testing.T) {
	repo := repository.NewInMemoryTaskRepository()
	svc := NewTaskService(repo)

	task, _ := svc.CreateTask(models.CreateTaskRequest{Title: "Test"})

	newDesc := "New description"
	req := models.UpdateTaskRequest{Description: &newDesc}

	updated, err := svc.UpdateTask(task.ID, req)
	if err != nil {
		t.Fatalf(msgExpectedNoError, err)
	}

	if updated.Description != newDesc {
		t.Errorf("expected description %s, got %s", newDesc, updated.Description)
	}
}

func TestTaskServiceCreateTaskRepositoryError(t *testing.T) {
	mockRepo := &repository.MockTaskRepository{
		CreateFunc: func(task *models.Task) error {
			return repository.ErrMockError
		},
	}
	svc := NewTaskService(mockRepo)

	req := models.CreateTaskRequest{Title: "Test"}
	_, err := svc.CreateTask(req)

	if err != repository.ErrMockError {
		t.Errorf("expected ErrMockError, got %v", err)
	}
}

func TestTaskServiceUpdateTaskRepositoryError(t *testing.T) {
	mockRepo := &repository.MockTaskRepository{
		GetByIDFunc: func(id string) (*models.Task, error) {
			return &models.Task{ID: "1", Title: "Test", Status: models.StatusTodo}, nil
		},
		UpdateFunc: func(task *models.Task) error {
			return repository.ErrMockError
		},
	}
	svc := NewTaskService(mockRepo)

	newTitle := "Updated"
	req := models.UpdateTaskRequest{Title: &newTitle}
	_, err := svc.UpdateTask("1", req)

	if err != repository.ErrMockError {
		t.Errorf("expected ErrMockError, got %v", err)
	}
}
