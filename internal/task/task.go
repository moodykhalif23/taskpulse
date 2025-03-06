package task

import (
	"time"
)

type Task struct {
	ID         string    `json:"id"`
	Type       string    `json:"type"`        // e.g., "email", "report"
	Payload    string    `json:"payload"`     // JSON-encoded data
	Schedule   string    `json:"schedule"`    // Cron expression (e.g., "*/5 * * * *")
	Priority   int       `json:"priority"`    // Higher = more urgent
	Retries    int       `json:"retries"`     // Max retry attempts
	RetryCount int       `json:"retry_count"` // Current attempts
	Status     string    `json:"status"`      // "pending", "running", "completed", "failed"
	CreatedAt  time.Time `json:"created_at"`
}
