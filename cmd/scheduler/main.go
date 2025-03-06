package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moodykhalif23/taskpulse/internal/api"
	"github.com/moodykhalif23/taskpulse/internal/queue"
	"github.com/moodykhalif23/taskpulse/internal/store"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

const APIKey = "super-secret-key" // In production, use env vars or a secret manager

func apiKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")
		if key != APIKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid API key"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r := gin.Default()
	rabbit, err := queue.NewRabbitMQ("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.Fatal("Failed to connect to RabbitMQ", zap.Error(err))
	}
	redisStore, err := store.NewRedisStore("localhost:6379", "", 0)
	if err != nil {
		logger.Fatal("Failed to connect to Redis", zap.Error(err))
	}
	scheduler := api.NewScheduler(rabbit, redisStore)

	r.GET("/tasks/list", apiKeyMiddleware(), scheduler.ListTasksHandler)
	r.POST("/tasks", scheduler.AddTaskHandler)
	r.GET("/metrics", gin.WrapH(promhttp.Handler())) // Prometheus metrics endpoint
	r.Static("/dashboard", "./web")

	scheduler.Start()
	logger.Info("Scheduler started on :8080")
	r.Run(":8080")
}
