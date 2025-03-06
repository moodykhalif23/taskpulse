# TaskPulse

TaskPulse is a distributed task scheduler built in Go. It efficiently manages scheduled tasks (e.g., sending emails, generating reports) using Cron-based scheduling, RabbitMQ for distribution, and Redis for job tracking—all while providing real-time monitoring via a dashboard and Prometheus metrics.

## Features

- **Scheduling:** Cron-based task scheduling.
- **Distributed Execution:** Workers receive tasks via RabbitMQ.
- **Resilience:** Configurable retries and task priorities.
- **Job Tracking:** Task status maintained in Redis.
- **Monitoring:** Real-time dashboard and Prometheus metrics.
- **Security:** API key authentication for task submissions.

## Tech Stack

- **Go** for concurrency and performance.
- **Gin** for the HTTP API & dashboard.
- **Redis** for persistent task metadata.
- **RabbitMQ** for messaging.
- **Prometheus** for metrics collection.
- **Zap** for structured logging.

## Prerequisites

- Go v1.21+
- Docker (for Redis, RabbitMQ, and Prometheus)
- Cross-platform (tested on Windows)

## Quickstart

### Clone & Build

```bash
git clone https://github.com/moodykhalif23/taskpulse.git
cd taskpulse
go mod tidy
go build -o taskpulse-scheduler ./cmd/scheduler
go build -o taskpulse-worker ./cmd/worker
Setup Dependencies (using Docker)
Redis:
docker run -d -p 6379:6379 redis
RabbitMQ:
docker run -d -p 5672:5672 rabbitmq:3
Prometheus:
Ensure a prometheus.yml is in the root, then run:
docker run -d -p 9090:9090 -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus

Start Scheduler:
./taskpulse-scheduler
Start Worker(s): (Run multiple for load balancing)
./taskpulse-worker

##Usage
Schedule a Task
Submit a task via the API (example using PowerShell):

Invoke-RestMethod -Uri "http://localhost:8080/tasks" -Method Post -Headers @{ "X-API-Key" = "super-secret-key"; "Content-Type" = "application/json" } -Body '{"type":"email","payload":"{\"to\":\"user@example.com\"}","schedule":"*/1 * * * *","priority":1,"retries":3}'
Monitor
Visit the Dashboard and Prometheus:

Dashboard: http://localhost:8080/dashboard
Prometheus: http://localhost:9090
## Contributing
  Fork the repository.
  Create your feature branch:
                             git checkout -b feature/your-feature
                             git commit -m "Your message"
                             Push and open a pull request.