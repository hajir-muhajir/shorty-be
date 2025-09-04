package repository

import (
	"context"
	"time"

	"github.com/hajir.muhajir/shorty-api/internal/domain"
)

type LinkRepository interface {
	Create(ctx context.Context, l *domain.Link) error
	FindByAlias(ctx context.Context, alias string) (*domain.Link, error)
	IncrementClick(ctx context.Context, linkID string) error
	SetUpdateAt(ctx context.Context, linkId string, t time.Time) error
}
