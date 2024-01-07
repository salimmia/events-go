package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(email string, userId int64) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email" : email,
		"user_id" : userId,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	err := godotenv.Load()
	if err != nil{
		return "", err
	}
	secretKey := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(secretKey))
}