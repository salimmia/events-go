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
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
	authenticated.POST("/users/logout", logout)
	server.GET("/users/refresh-token", RefreshToken)

	server.POST("/users/signup", SingUp)
	server.POST("/users/login", LogIn)
	server.GET("/users", GetUsers)
   	server.GET("/users/:id", GetUser)
}