package queue

import (
	"context"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/ljcnh/task-center/internal/model"
)

type TaskHandler interface {
	HandleTask(ctx context.Context, task *model.Task) error
}

type RocketMQConsumer struct {
	consumer rocketmq.PushConsumer
	handler  TaskHandler
}

func NewRocketMQConsumer(handler TaskHandler) (*RocketMQConsumer, error) {
	return &RocketMQConsumer{
		handler: handler,
	}, nil
}

func (c *RocketMQConsumer) Close() error {
	return c.consumer.Shutdown()
}
