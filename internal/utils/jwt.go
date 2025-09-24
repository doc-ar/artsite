package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CreateJWT generates a signed JWT for the given username
func CreateJWT(username string) (string, error) {
	secretKey := []byte(os.Getenv("TOKEN_SECRET"))
	if len(secretKey) == 0 {
		return "", fmt.Errorf("TOKEN_SECRET not set")
	}

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(20 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// VerifyJWT parses and validates a JWT token string
func VerifyJWT(tokenString string) (*jwt.Token, error) {
	secretKey := []byte(os.Getenv("TOKEN_SECRET"))
	if len(secretKey) == 0 {
		return nil, fmt.Errorf("TOKEN_SECRET not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
