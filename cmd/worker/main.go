package main

import (
	"encoding/json"

	"github.com/moodykhalif23/taskpulse/internal/queue"
	"github.com/moodykhalif23/taskpulse/internal/store"
	"github.com/moodykhalif23/taskpulse/internal/task"
	"github.com/moodykhalif23/taskpulse/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	logger.Init()
	defer logger.Sync()

	rabbit, err := queue.NewRabbitMQ("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.Logger.Fatal("Failed to connect to RabbitMQ", zap.Error(err))
	}
	redisStore, err := store.NewRedisStore("localhost:6379", "", 0)
	if err != nil {
		logger.Logger.Fatal("Failed to connect to Redis", zap.Error(err))
	}

	msgs, err := rabbit.Consume()
	if err != nil {
		logger.Logger.Fatal("Failed to consume messages", zap.Error(err))
	}

	for msg := range msgs {
		var t task.Task
		json.Unmarshal(msg.Body, &t)

		t.Status = "running"
		redisStore.UpdateTask(&t)

		logger.Logger.Info("Processing task", zap.String("task_id", t.ID), zap.String("type", t.Type))
		success := executeTask(t)

		if !success && t.RetryCount < t.Retries {
			t.RetryCount++
			t.Status = "pending"
			redisStore.UpdateTask(&t)
			rabbit.PublishTask(t)
			logger.Logger.Warn("Task retry scheduled", zap.String("task_id", t.ID), zap.Int("retry_count", t.RetryCount))
		} else if success {
			t.Status = "completed"
			redisStore.UpdateTask(&t)
			logger.Logger.Info("Task completed", zap.String("task_id", t.ID))
		} else {
			t.Status = "failed"
			redisStore.UpdateTask(&t)
			logger.Logger.Error("Task failed", zap.String("task_id", t.ID))
		}
		msg.Ack(false)
	}
}

func executeTask(t task.Task) bool {
	return t.Type == "email" // Simulate success
}
