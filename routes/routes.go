package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/salimmia/events-go/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)

	server.POST("/signup", SingUp)
	server.POST("/login", LogIn)
}