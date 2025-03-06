package task_test

import (
	"testing"
	"time"

	"github.com/moodykhalif23/taskpulse/internal/task"
)

func TestTaskCreation(t *testing.T) {
	task := task.Task{
		ID:        "test-id",
		Type:      "email",
		Payload:   `{"to": "test@example.com"}`,
		Schedule:  "*/5 * * * *",
		Priority:  1,
		Retries:   3,
		CreatedAt: time.Now(),
	}

	if task.ID != "test-id" {
		t.Errorf("expected ID 'test-id', got %s", task.ID)
	}
	if task.Status != "" {
		t.Errorf("expected empty status, got %s", task.Status)
	}
}
