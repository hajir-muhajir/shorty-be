package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/hajir.muhajir/shorty-be/internal/domain"
	"github.com/hajir.muhajir/shorty-be/internal/repository"
)

type (
	PasswordHasher interface {
		Hash(string) (string, error)
		Verify(hash, pw string) bool
	}

	TokenSigner interface {
		Sign(userID string) (string, error)
	}
)

type AuthUC struct {
	users repository.UserRepository
	hash  PasswordHasher
	jwt   TokenSigner
}

func NewAuthUC(users repository.UserRepository, hash PasswordHasher, jwt TokenSigner) *AuthUC {
	return &AuthUC{
		users: users,
		hash:  hash,
		jwt:   jwt,
	}
}

func (uc *AuthUC) Register(ctx context.Context, email, password string) (string, error) {
	if _, err := uc.users.FindByEmail(ctx, email); err == nil {
		return "", ErrConflict
	}

	h, err := uc.hash.Hash(password)
	if err != nil {
		return "", nil
	}

	u := &domain.User{
		ID:           uuid.NewString(),
		Email:        email,
		PasswordHash: h,
	}

	if err := uc.users.Create(ctx, u); err != nil {
		return "", err
	}
	return uc.jwt.Sign(u.ID)
}

func (uc *AuthUC) Login(ctx context.Context, email, password string) (string, error) {
	u, err := uc.users.FindByEmail(ctx, email)
	if err != nil {
		return "", ErrUnauthorized
	}

	if !uc.hash.Verify(u.PasswordHash, password) {
		return "", ErrUnauthorized
	}

	return uc.jwt.Sign(u.ID)
}
