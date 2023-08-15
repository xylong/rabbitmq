package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	// 连接mq
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", "root", "123456", "localhost", 5672)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatal(err)
	}
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
	body := "Hello World!"
	if err := c.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	}); err != nil {
		log.Fatal(err)
	}

	log.Printf(" [x] Sent %s\n", body)
}
