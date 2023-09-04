package Lib

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/AppInit"
	"rabbitmq/constant"
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

// SendMessage 发送消息
func (m *MQ) SendMessage(queue string, message string) error {
	// 用户注册队列
	q1, err := m.Channel.QueueDeclare(queue, false, false, false, false, nil)
	if err != nil {
		return err
	}

	// 用户注册通知队列
	q2, err := m.Channel.QueueDeclare(queue+".notify", false, false, false, false, nil)
	if err != nil {
		return err
	}

	// 申明交换机
	err = m.Channel.ExchangeDeclare(constant.UserExchange, "direct", false, false, false, false, nil)
	if err != nil {
		return err
	}

	err = m.Channel.QueueBind(q1.Name, constant.UserRegister, constant.UserExchange, false, nil)
	if err != nil {
		return err
	}
	err = m.Channel.QueueBind(q2.Name, constant.UserRegister, constant.UserExchange, false, nil)
	if err != nil {
		return err
	}

	return m.Channel.Publish(constant.UserExchange, constant.UserRegister, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
}
