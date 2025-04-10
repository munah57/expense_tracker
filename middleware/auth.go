package middleware

import (
	"time"
	"fmt"
	"net/http"
	"strings"
	"os"
	jwt "github.com/golang-jwt/jwt/v4"

)
//generate JWT token by taking in UserID
func GenerateJWT(userID uint) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret)) //KEY 
	if err != nil {
		return "", err
	}

	return tokenString, nil 
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method")
	} //check signing method
	return []byte(secret), nil
	})
	
	if err != nil {
		return nil, err
	}
	return token, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "authorization header required", http.StatusUnauthorized)
			return 
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		_, err := VerifyJWT(tokenString)
		if err != nil {
			http.Error(w, "invalid toke", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})

}

func GetUserIDFromToken (r *http.Request) (uint, error) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			return 0, fmt.Errorf("authorization header required")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := VerifyJWT(tokenString)
		if err != nil {
			return 0, fmt.Errorf("invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return 0, fmt.Errorf("could not parse claims")
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			return 0, fmt.Errorf("user id not found in claims")
		}

			return uint(userID), nil

	}


