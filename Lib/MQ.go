package Lib

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/AppInit"
)

const (
	Register = "register"
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

func (m *MQ) SendMessage(queue string, message string) error {
	q, err := m.Channel.QueueDeclare(queue, false, false, false, false, nil)
	if err != nil {
		return err
	}

	return m.Channel.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
}
