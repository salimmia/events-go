package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salimmia/events-go/models"
)

func SingUp(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Wrong input given"})
		return
	}

	u, _ := models.GetUserByEmail(user.Email)
	if u != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Already registered"})
		return
	}

	err = user.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Unsuccessful registration"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message" : "Successfully Registered", "user": user})
}

func LogIn(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Could not fetch user login info."})
		return
	}

	err = user.ValidateCredentials()
	// log.Println(err)

	if err != nil{
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message" : "Successfully logged In"})
}