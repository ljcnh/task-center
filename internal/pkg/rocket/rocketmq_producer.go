package queue

// RocketMQProducer RocketMQ生产者
type RocketMQProducer struct {
}

// NewRocketMQProducer 创建RocketMQ生产者
func NewRocketMQProducer() (*RocketMQProducer, error) {
	return &RocketMQProducer{}, nil
}
