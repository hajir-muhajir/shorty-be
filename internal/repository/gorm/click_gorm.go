package gormrepo

import (
	"context"

	"github.com/hajir.muhajir/shorty-api/internal/domain"
	"github.com/hajir.muhajir/shorty-api/internal/repository"
	"gorm.io/gorm"
)

type clickGorm struct {
	db *gorm.DB
}

func NewClickGorm(db *gorm.DB) repository.ClickRepository {
	return &clickGorm{db: db}
}

func (r *clickGorm) Insert(ctx context.Context, c *domain.Click) error {
	return r.db.WithContext(ctx).Create(c).Error
}
