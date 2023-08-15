package main

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/AppInit"
)

func main() {
	// 连接mq
	conn := AppInit.GetConn()
	defer conn.Close()

	// 获取channel
	c, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// 创建队列
	q, err := c.QueueDeclare("test", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 发送
	body := "test001"
	if err := c.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	}); err != nil {
		log.Fatal(err)
	}

	log.Printf(" [x] Sent %s\n", body)
}
