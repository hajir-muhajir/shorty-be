package service

import (
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestJWTSigner_SignAndParse_Valid(t *testing.T) {
	secret := "test-secret"
	ttl := time.Hour
	signer := NewJWTSigner(secret, ttl)

	tokenStr, err := signer.Sign("user-123")
	if err != nil {
		t.Fatalf("sign error %v", err)
	}

	claims := &jwt.RegisteredClaims{}
	tok, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})

	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if !tok.Valid {
		t.Fatalf("expected token valid")
	}
	if claims.Subject != "user-123" {
		t.Fatalf("expected user-123, got %s", claims.Subject)
	}
	if claims.ExpiresAt == nil {
		t.Fatalf("expected exp set")
	}
}

func TestJWTSigner_InvalidSignature(t *testing.T) {
	signer := NewJWTSigner("secret-a", time.Hour)
	tokenStr, err := signer.Sign("user-123")
	if err != nil {
		t.Fatalf("Sign error %v", err)
	}

	claims := &jwt.RegisteredClaims{}
	tok, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		return []byte("secret-b"), nil
	})

	if err == nil || tok.Valid {
		t.Fatalf("expected invalid token due to bad signature")
	}
}

func TestJWTSigner_Expired(t *testing.T) {
	secret := "secret"
	signer := NewJWTSigner(secret, -1*time.Second)
	tokenStr, err := signer.Sign("user-123")
	if err != nil {
		t.Fatalf("sign error: %v", err)
	}

	claims := &jwt.RegisteredClaims{}
	tok, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err == nil || tok.Valid {
		t.Fatalf("expected expired token, got valid")
	}
	if !strings.Contains(strings.ToLower(err.Error()), "expired") {
		t.Fatalf("expected error mentions expired, got: %v", err)
	}
	if claims.ExpiresAt == nil || claims.ExpiresAt.Time.After(time.Now()) {
		t.Fatalf("expected exp in the past, got %v", claims.ExpiresAt)
	}
}
