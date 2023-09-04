package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rabbitmq/Lib"
	"rabbitmq/UserReg/Models"
	"rabbitmq/constant"
	"strconv"
	"time"
)

func main() {
	router := gin.Default()

	router.Handle(http.MethodPost, "user", func(context *gin.Context) {
		userModel := Models.NewUserModel()
		if err := context.BindJSON(&userModel); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"result": "param error"})
		} else {
			// 模拟入库
			userModel.UserID = int(time.Now().Unix())
			mq := Lib.NewMQ()
			err = mq.SendMessage(constant.UserExchange, constant.UserRegister, strconv.Itoa(userModel.UserID))
			defer mq.Channel.Close()
			if err != nil {
				log.Println(err)
			}

			context.JSON(http.StatusOK, gin.H{"result": userModel})
		}
	})

	c := make(chan error)

	go func() {
		if err := router.Run(":8080"); err != nil {
			c <- err
		}
	}()

	go func() {
		if err := Lib.UserQueueInit(); err != nil {
			c <- err
		}
	}()

	err := <-c
	log.Fatal(err)
}
