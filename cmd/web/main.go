package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/salimmia/events-go/db"
	"github.com/salimmia/events-go/models"
)

func main() {
	server := gin.Default()
	db.InitDB()

	err := godotenv.Load()

	if err != nil{
		log.Println("Error loading .env file")
		return
	}

	PORT := os.Getenv("PORT")

	log.Println(PORT)

	server.GET("/events", getEvents)

	server.POST("/events", createEvent)

	server.Run(":" + PORT)
}

func getEvents(context *gin.Context){
	events := models.GetEvent()
	context.JSON(http.StatusOK, gin.H{"message": "Hello!", "events" : events})
}

func createEvent(context *gin.Context){
	event := models.Event{}

	err := context.ShouldBindJSON(&event)

	log.Println(event)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : "don't create any event"})
		return
	}

	event.ID = 1
	event.UserId = 1;

	event.DateTime = time.Now()

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "created event Successfully", "event" : event})
}