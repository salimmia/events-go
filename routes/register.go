package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/salimmia/events-go/models"
)

func registerForEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not parse eventId"})
		return
	}

	userId := context.GetInt64("user_id")

	event, err := models.GetEventById(eventId)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Could not fetch event"})
		return
	}

	isRegistered, err := event.IsAlreadyRegistered(userId)
	if err != nil && isRegistered{
		context.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	registerId, err := event.Register(userId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Could not register user for this event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message" : "Successfully registered!", "registerId" : registerId})
}

func cancelRegistration(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not parse eventId"})
		return
	}

	userId := context.GetInt64("user_id")

	event, err := models.GetEventById(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Could not fetch event."})
		return
	}

	err = event.CancelRegistration(userId)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message" : err.Error()})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"message" : "Cancelled Registration"})
}