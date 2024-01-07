package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/salimmia/events-go/db"
	"github.com/salimmia/events-go/routes"
)

func main() {
	db.InitDB()
	
	err := godotenv.Load()
	if err != nil{
		log.Println("Error loading .env file")
		return
	}
	PORT := os.Getenv("PORT")

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":" + PORT)
}
