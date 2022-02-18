package router

import (
	"chattsimulator/chatt/controller"
	"chattsimulator/chatt/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.Static("/asset", "./asset")
	r.Static("/public", "./public")
	r.LoadHTMLGlob("./pages/**/*")
	r.NoRoute(error404)
	r.NoMethod(error404)

	user := r.Group("/user")
	user.POST("/register", controller.RegisterUser)
	userMessage := user.Group("/room")
	userMessage.Use(middleware.User)
	userMessage.GET("/my", controller.MyRoom)
	userMessage.POST("/create", controller.CreateRoom)
	userMessage.POST("/create/group", controller.CreateRoomGroup)
	userMessage.GET("/kick/user/:room_id/:nohp", controller.KickFromRoom)
	userMessage.GET("/members/:room_id", controller.RoomMember)
	userMessage.POST("/message/sent", controller.SentMessage)
	userMessage.GET("/message/:room_id", controller.MyRoomMessage)
}

func error404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error/404", gin.H{
		"title": "Error 404",
	})
}
