package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	TaskDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "task_duration_seconds",
		Help: "Duration of task execution",
	})
	TaskFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "task_failures_total",
		Help: "Total number of failed tasks",
	})
)

func Init() {
	prometheus.MustRegister(TaskDuration, TaskFailures)
}

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	}
}
