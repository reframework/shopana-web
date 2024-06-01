package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserIdClaims struct {
	UserID uuid.UUID `json:"alpha"`
	jwt.RegisteredClaims
}

type TokenString = string

func CreateAuthToken(userID uuid.UUID, salt string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserIdClaims{
		UserID: userID,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		},
	})

	tokenString, err := token.SignedString([]byte(salt))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateAuthToken(tokenStr string, salt string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserIdClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(salt), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*UserIdClaims)
	if ok && token.Valid && claims.UserID != uuid.Nil {
		return claims.UserID, nil
	}

	return uuid.Nil, fmt.Errorf("invalid token")
}
