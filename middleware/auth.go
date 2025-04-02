package middleware

import (
	"time"
	"fmt"
	"net/http"
	"os"
	jwt "github.com/golang-jwt/jwt/v4"

)
//generate JWT token 
func GenerateJWT(userID uint) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err :=token.SignedString([]byte(secret))
	if err != nil {
		return "", nil
	}

	return tokenString, nil 
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error)) {

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method")
	}
	return []byte(secret), nil

}
