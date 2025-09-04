package repository

import (
	"context"

	"github.com/hajir.muhajir/shorty-be/internal/domain"
)

type UserRepository interface{
	Create(ctx context.Context, u *domain.User)error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}