package gormrepo

import (
	"context"

	"github.com/hajir.muhajir/shorty-be/internal/domain"
	"github.com/hajir.muhajir/shorty-be/internal/repository"
	"gorm.io/gorm"
)

type userGorm struct {
	db *gorm.DB
}

func NewUserGorm(db *gorm.DB) repository.UserRepository {
	return &userGorm{
		db: db,
	}
}

func (r *userGorm) Create(ctx context.Context, u *domain.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}
func (r *userGorm) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var u domain.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &u, nil
}
