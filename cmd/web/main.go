package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/salimmia/events-go/db"
	"github.com/salimmia/events-go/models"
)

func main() {
	db.InitDB()

	server := gin.Default()
	
	err := godotenv.Load()

	if err != nil{
		log.Println("Error loading .env file")
		return
	}

	PORT := os.Getenv("PORT")

	log.Println(PORT)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	server.POST("/events", createEvent)

	server.Run(":" + PORT)
}

func getEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id."})
		return;
	}

	context.JSON(http.StatusOK, gin.H{"message" : "Successfully find this event", "event" : event})
}

func getEvents(context *gin.Context){
	events, err := models.GetEvents()

	if err != nil{
		log.Println(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Hello!", "events" : events})
}

func createEvent(context *gin.Context){
	event := models.Event{}

	err := context.ShouldBindJSON(&event)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : "don't create any event"})
		return
	}

	event.DateTime = time.Now()
	event.UserId = 1

	err = event.Save()

	if err != nil{
		log.Println("error  ", err)
		return
	}

	log.Println(event)

	context.JSON(http.StatusCreated, gin.H{"message": "created event Successfully", "event" : event})
}