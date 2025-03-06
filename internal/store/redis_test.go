package store_test

import (
	"testing"
	"time"

	"github.com/moodykhalif23/taskpulse/internal/store"
	"github.com/moodykhalif23/taskpulse/internal/task"
)

func TestRedisStore(t *testing.T) {
	s, err := store.NewRedisStore("localhost:6379", "", 0)
	if err != nil {
		t.Fatalf("failed to connect to Redis: %v", err)
	}

	task := &task.Task{
		ID:        "test-123",
		Type:      "report",
		Payload:   `{"data": "test"}`,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	// Test SaveTask
	if err := s.SaveTask(task); err != nil {
		t.Errorf("failed to save task: %v", err)
	}

	// Test GetTask
	retrieved, err := s.GetTask("test-123")
	if err != nil {
		t.Errorf("failed to get task: %v", err)
	}
	if retrieved.Type != "report" {
		t.Errorf("expected type 'report', got %s", retrieved.Type)
	}
}
