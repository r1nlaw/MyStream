package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Maker interface {
	CreateToken(userID int64) (string, error)
	VerifyToken(token string) (*jwt.Token, error)
}

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) *JWTMaker {
	return &JWTMaker{secretKey: secretKey}
}

func (j *JWTMaker) CreateToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userID,
		"expired_at": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", fmt.Errorf("could not sign token: %v", err)
	}
	return signedToken, nil
}

func (j *JWTMaker) VerifyToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not parse token: %v", err)
	}

	return token, nil
}
