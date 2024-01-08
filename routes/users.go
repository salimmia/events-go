package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salimmia/events-go/models"
	"github.com/salimmia/events-go/utils"
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
	if err != nil{
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Could not generate jwt token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message" : "Successfully logged In", "token" : token})
}