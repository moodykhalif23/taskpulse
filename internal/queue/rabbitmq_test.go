package queue_test

import (
	"testing"

	"github.com/moodykhalif23/taskpulse/internal/queue"
	"github.com/moodykhalif23/taskpulse/internal/task"
)

func TestRabbitMQPublish(t *testing.T) {
	q, err := queue.NewRabbitMQ("amqp://guest:guest@localhost:5672/")
	if err != nil {
		t.Fatalf("failed to connect to RabbitMQ: %v", err)
	}

	task := task.Task{
		ID:      "test-queue",
		Type:    "email",
		Payload: `{"to": "test@example.com"}`,
	}

	if err := q.PublishTask(task); err != nil {
		t.Errorf("failed to publish task: %v", err)
	}
}
