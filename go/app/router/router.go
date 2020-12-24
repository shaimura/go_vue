package router

import (
	"../controller"
	"../db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	router.Use(cors.New(config))

	userhandler := controller.UserHandler{
		db.Get(),
	}

	chathandler := controller.ChatHandler{
		db.Get(),
	}

	router.POST("/usersinup", userhandler.SignUp)
	router.POST("/usersinin", userhandler.SignIn)
	router.POST("/getusers", userhandler.GetUsers)
	router.POST("/userchatroom", userhandler.UserChatRoom)
	router.POST("/sendusermessage", chathandler.SendUserMessage)
	router.GET("/getusermessage/:id", chathandler.GetUserMessage)
	router.GET("/ws", controller.Chat)
	go controller.HandleMessages()

	router.Run(":8888")
}
