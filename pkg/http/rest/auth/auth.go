package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AppClaims struct {
	ID uint `json:"id,omitempty"`
	//Roles []string `json:"roles,omitempty"`
	jwt.StandardClaims
}

type Authenticator interface {
	ParseToken(token string) (*AppClaims, error)
	GenerateToken(id uint) (string, error)
}

type authenticator struct {
	secretKey []byte
}

func NewAuthenticator(secretKey string) Authenticator {
	return &authenticator{[]byte(secretKey)}
}

func (a authenticator) ParseToken(tokenString string) (*AppClaims, error) {
	var claims AppClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return a.secretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with claims")
	}

	if claim, ok := token.Claims.(*AppClaims); ok && token.Valid {
		return claim, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func (a authenticator) GenerateToken(id uint) (string, error) {
	var claims AppClaims
	claims.ID = id
	claims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.secretKey)
}

func (a authenticator) Validate(claims AppClaims) error {
	if claims.ID == 0 {
		return fmt.Errorf("user ID must be set")
	}
	return nil
}
