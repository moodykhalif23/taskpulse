package queue

import (
	"encoding/json"

	"github.com/moodykhalif23/taskpulse/internal/task"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	_, err = ch.QueueDeclare("tasks", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return &RabbitMQ{conn: conn, channel: ch}, nil
}

func (r *RabbitMQ) PublishTask(t task.Task) error {
	body, _ := json.Marshal(t)
	return r.channel.Publish("", "tasks", false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        body,
		Priority:    uint8(t.Priority),
	})
}

func (r *RabbitMQ) Consume() (<-chan amqp091.Delivery, error) {
	return r.channel.Consume("tasks", "", false, false, false, false, nil)
}
