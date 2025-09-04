package gormrepo

import (
	"context"
	"time"

	"github.com/hajir.muhajir/shorty-be/internal/domain"
	"github.com/hajir.muhajir/shorty-be/internal/repository"
	"gorm.io/gorm"
)

type linkGorm struct {
	db *gorm.DB
}

func NewLinkGorm(db *gorm.DB) repository.LinkRepository {
	return &linkGorm{db: db}
}

func (r *linkGorm) Create(ctx context.Context, l *domain.Link) error {
	return r.db.WithContext(ctx).Create(l).Error
}

func (r *linkGorm) FindByAlias(ctx context.Context, alias string) (*domain.Link, error) {
	var link domain.Link
	if err := r.db.WithContext(ctx).Where("alias = ?", alias).First(&link).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &link, nil
}

func (r *linkGorm) IncrementClick(ctx context.Context, linkID string) error {
	return r.db.WithContext(ctx).
		Model(&domain.Link{}).
		Where("id = ?", linkID).
		UpdateColumn("click_count", gorm.Expr("click_count + 1")).Error

}
func (r *linkGorm) SetUpdateAt(ctx context.Context, linkId string, t time.Time) error {
	return r.db.WithContext(ctx).
		Model(&domain.Link{}).
		Where("id = ?", linkId).
		UpdateColumn("updated_at", t).Error
}
