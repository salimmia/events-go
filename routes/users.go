package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/salimmia/events-go/helpers"
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

	accessToken, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Could not generate jwt token"})
		return
	}

	refreshToken, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Could not generate jwt token"})
		return
	}

	helpers.SetCookie(context, "jwt", accessToken, time.Now().Add(time.Hour*1))
    helpers.SetCookie(context, "refresh_token", refreshToken, time.Now().Add(time.Hour*24*7))

	context.JSON(http.StatusOK, gin.H{"message" : "Successfully logged In", "access" : accessToken, "refresh_token" : refreshToken})
}

func RefreshToken(context *gin.Context){
	refreshToken, err := context.Cookie("refresh_token")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "missing refresh token"})
		return
	}

	userId, email, err := utils.VerifyToken(refreshToken) 
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}

	user := models.User{
		ID: userId,
		Email: email,
	}

	//    Db.Where("ID = ?", claims.UserID).First(&user)

	accessToken, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate access token"})
		return
	}
	helpers.SetCookie(context, "jwt", accessToken, time.Now().Add(time.Hour*1))

	context.JSON(http.StatusOK, gin.H{"message": "access token refreshed successfully"})
}

func logout(context *gin.Context){
	helpers.ClearCookie(context, "jwt")
	helpers.ClearCookie(context, "refresh_token")

	context.JSON(http.StatusOK, gin.H{
		"message": "User logged out successfully",
	})
}


func GetUsers(context *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error while querying database"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Users" : users})
}

func GetUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not parse event id."})
		return
	}

	user, err := models.GetUserByID(userId)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error while querying database"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"User" : user})
}
