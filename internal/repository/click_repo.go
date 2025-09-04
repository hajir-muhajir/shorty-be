package repository

import (
	"context"

	"github.com/hajir.muhajir/shorty-be/internal/domain"
)

type ClickRepository interface {
	Insert(ctx context.Context, c *domain.Click) error
}
