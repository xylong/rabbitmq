package Lib

import (
	"fmt"
	"rabbitmq/constant"
)

// UserQueueInit 初始化用户相关队列
func UserQueueInit() error {
	mq := NewMQ()
	if mq == nil {
		return fmt.Errorf("mq init error")
	}
	defer mq.Channel.Close()

	// 申明交换机
	err := mq.Channel.ExchangeDeclare(constant.UserExchange, "direct", false, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("declare exchange error:%s", err)
	}

	err = mq.DeclareQueueAndBind(constant.UserExchange, constant.UserRegister, constant.RegisterQueue, constant.RegisterNotifyQueue)
	if err != nil {
		return fmt.Errorf("queue bind error:%s", err)
	}

	return nil
}
