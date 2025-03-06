FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o taskpulse-scheduler ./cmd/scheduler
RUN go build -o taskpulse-worker ./cmd/worker

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/taskpulse-scheduler .
COPY --from=builder /app/taskpulse-worker .
CMD ["./taskpulse-scheduler"] # Or "./taskpulse-worker"