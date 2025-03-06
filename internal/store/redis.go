package store

import (
	"context"
	"encoding/json"

	"github.com/moodykhalif23/taskpulse/internal/task"
	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(addr, password string, db int) (*RedisStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisStore{client: client}, nil
}

func (r *RedisStore) SaveTask(t *task.Task) error {
	ctx := context.Background()
	data, _ := json.Marshal(t)
	return r.client.Set(ctx, t.ID, data, 0).Err()
}

func (r *RedisStore) GetTask(id string) (*task.Task, error) {
	ctx := context.Background()
	data, err := r.client.Get(ctx, id).Bytes()
	if err != nil {
		return nil, err
	}
	var t task.Task
	json.Unmarshal(data, &t)
	return &t, nil
}

func (r *RedisStore) UpdateTask(t *task.Task) error {
	return r.SaveTask(t) // Overwrites existing task
}

func (r *RedisStore) ListTasks() ([]*task.Task, error) {
	ctx := context.Background()
	keys, err := r.client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	var tasks []*task.Task
	for _, key := range keys {
		t, err := r.GetTask(key)
		if err != nil {
			continue // Skip invalid entries
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
