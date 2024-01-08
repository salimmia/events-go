package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	server.POST("/events", createEvent)

	server.PUT("/events/:id", UpdateEvent)
	server.DELETE("/events/:id", DeleteEvent)
	server.POST("/signup", SingUp)
	server.POST("/login", LogIn)
}