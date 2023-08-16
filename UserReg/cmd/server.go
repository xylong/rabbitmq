package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rabbitmq/Lib"
	"rabbitmq/UserReg/Models"
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
			if err := mq.SendMessage(Lib.Register, strconv.Itoa(userModel.UserID)); err != nil {
				log.Println(err)
			}
			context.JSON(http.StatusOK, gin.H{"result": userModel})
		}
	})

	router.Run(":8080")
}
