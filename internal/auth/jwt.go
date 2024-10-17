package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hazaloolu/openUp_backend/internal/model"
)

var secretKey = GenerateSecureKey()

type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"userID"`
	jwt.RegisteredClaims
}

func GenerateJwt(user *model.User, secretKey string) (string, error) {
	claims := &Claims{
		Username: user.Username,
		UserID:   user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJwt(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		return []byte(secretKey), nil

	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")

	}

	return claims, nil
}
