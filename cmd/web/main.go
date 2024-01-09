package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salimmia/events-go/db"
	"github.com/salimmia/events-go/helpers"
	"github.com/salimmia/events-go/routes"
)

func main() {
	helpers.LoadConfig(".env")
	db.InitDB()

	appConfig := helpers.AppConfig

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":" + appConfig.PORT)
}
