package AppInit

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var MQ *amqp.Connection

func init() {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", "root", "123456", "localhost", 5672)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatal(err)
	}
	MQ = conn
}

func GetConn() *amqp.Connection {
	return MQ
}
