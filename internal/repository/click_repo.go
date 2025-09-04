package repository

import (
	"context"

	"github.com/hajir.muhajir/shorty-api/internal/domain"
)

type ClickRepository interface {
	Insert(ctx context.Context, c *domain.Click) error
}
