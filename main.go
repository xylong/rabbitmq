package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", "root", "123456", "localhost", 5672)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}
