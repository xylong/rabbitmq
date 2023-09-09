package main

import (
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"rabbitmq/Lib"
	"rabbitmq/constant"
	"time"
)

// SendEmail 发邮件
func SendEmail(message <-chan amqp.Delivery, consumer string) {
	for msg := range message {
		fmt.Printf("%s send email to user:%s\n", consumer, string(msg.Body))

		// 模拟c1出问题
		if consumer == "c1" {
			msg.Reject(true) // 重新入列
			continue
		}

		time.Sleep(time.Second * 1) // 模拟处理消息
		msg.Ack(false)
	}
}

func main() {
	var c *string
	c = flag.String("c", "", "消费者名称")
	flag.Parse()

	if *c == "" {
		log.Fatalln("消费者名称不能为空")
	}

	mq := Lib.NewMQ()
	mq.Consume(constant.UserRegister, *c, SendEmail)
	defer mq.Channel.Close()
}
