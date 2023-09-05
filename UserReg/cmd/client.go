package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"rabbitmq/Lib"
	"rabbitmq/constant"
	"time"
)

func SendEmail(message <-chan amqp.Delivery) {
	for msg := range message {
		fmt.Println(map[string]any{
			"id":   msg.MessageId,
			"tag":  msg.DeliveryTag,
			"body": string(msg.Body),
		})

		time.Sleep(time.Second * 1) // 模拟处理消息
		msg.Ack(false)
	}
}

func main() {
	mq := Lib.NewMQ()
	mq.Consume(constant.UserRegister, "c1", SendEmail)
	defer mq.Channel.Close()
}
