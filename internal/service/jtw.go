package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTSigner struct{
	secret []byte
	ttl time.Duration
}

func NewJWTSigner(secret string, ttl time.Duration) *JWTSigner{
	return &JWTSigner{
		secret: []byte(secret),
		ttl: ttl,
	}
}

func (j *JWTSigner) Sign(userID string)(string, error){
	claims := jwt.MapClaims{
		"sub": userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(j.ttl).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(j.secret)
}