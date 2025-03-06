package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moodykhalif23/taskpulse/internal/queue"
	"github.com/moodykhalif23/taskpulse/internal/store"
	"github.com/moodykhalif23/taskpulse/internal/task"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron  *cron.Cron
	queue *queue.RabbitMQ
	store *store.RedisStore
}

func NewScheduler(q *queue.RabbitMQ, s *store.RedisStore) *Scheduler {
	return &Scheduler{
		cron:  cron.New(),
		queue: q,
		store: s,
	}
}

func (s *Scheduler) AddTaskHandler(c *gin.Context) {
	var t task.Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	t.ID = uuid.New().String()
	t.Status = "pending"
	t.CreatedAt = time.Now()

	// Store task metadata
	if err := s.store.SaveTask(&t); err != nil {
		c.JSON(500, gin.H{"error": "failed to save task"})
		return
	}

	// Schedule with Cron
	_, err := s.cron.AddFunc(t.Schedule, func() {
		s.queue.PublishTask(t)
	})
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid schedule"})
		return
	}

	c.JSON(200, gin.H{"message": "task scheduled", "task_id": t.ID})
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) ListTasksHandler(c *gin.Context) {
	tasks, err := s.store.ListTasks() // We'll implement this
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to list tasks"})
		return
	}
	c.JSON(200, tasks)
}
