package controllers

import (
	"time"

	"github.com/RND2002/goChatApp/models"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserID   uint   `json:"user_id"`
	jwt.RegisteredClaims
}

var secret string = "secretKey"

func GenerateToken(requestedUser models.User) (string, error) {
	claims := CustomClaims{
		Username:         requestedUser.Username,
		Password:         requestedUser.Password,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24))},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}
	return t, nil
}

func ValidateToken(tokenString string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok {
		return nil, err
	}

	return claims, err
}

func RefreshToken(tokenString string) (string, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func AuthenticateToken(tokenString string) (bool, error) {
	_, err := ValidateToken(tokenString)
	if err != nil {
		return false, err
	}
	return true, nil
}
