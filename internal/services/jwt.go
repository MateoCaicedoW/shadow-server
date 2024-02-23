package services

import (
	"errors"
	"time"

	jwtToken "github.com/golang-jwt/jwt/v5"
	"github.com/leapkit/core/envor"

	"github.com/shadow/backend/internal/models"
)

type JWTService interface {
	// Generate a new JWT token
	GenerateToken(user models.User) (string, error)

	// Validate a JWT token
	ValidateToken(token string) (map[string]interface{}, error)
}

var _ JWTService = (*jwt)(nil)

type jwt struct{}

var JWTSecret = []byte(envor.Get("JWT_SECRET", "secret"))

func JWT() *jwt {
	return &jwt{}
}

func (j *jwt) GenerateToken(user models.User) (string, error) {
	token := jwtToken.NewWithClaims(jwtToken.SigningMethodHS256,
		jwtToken.MapClaims{
			"user": user,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func (j *jwt) ValidateToken(token string) (map[string]interface{}, error) {
	t, err := jwtToken.Parse(token, func(t *jwtToken.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwtToken.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims["user"].(map[string]interface{}), nil
}
