package Lib

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/AppInit"
)

type MQ struct {
	*amqp.Channel
}

func NewMQ() *MQ {
	c, err := AppInit.GetConn().Channel()
	if err != nil {
		log.Println(err)
		return nil
	}

	return &MQ{Channel: c}
}

// DeclareQueueAndBind 申明队列并且绑定交换机路由
func (m *MQ) DeclareQueueAndBind(exchange, key string, queues ...string) error {
	for _, queue := range queues {
		q, err := m.Channel.QueueDeclare(queue, false, false, false, false, nil)
		if err != nil {
			return err
		}

		err = m.Channel.QueueBind(q.Name, key, exchange, false, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

// SendMessage 发送消息
func (m *MQ) SendMessage(exchange, key string, message string) error {
	return m.Channel.Publish(exchange, key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
}
