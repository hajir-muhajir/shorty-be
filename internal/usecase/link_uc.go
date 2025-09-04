package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/hajir.muhajir/shorty-api/internal/domain"
	"github.com/hajir.muhajir/shorty-api/internal/repository"
	"github.com/hajir.muhajir/shorty-api/internal/service"
)

type LinkUC struct {
	repo  repository.LinkRepository
	alias *service.AliasGenerator
	now   func() time.Time
}

func NewLinkUC(r repository.LinkRepository, a *service.AliasGenerator) *LinkUC {
	return &LinkUC{
		repo:  r,
		alias: a,
		now:   time.Now,
	}
}

type CreateLinkRequest struct {
	UserID      string
	OriginalURL string
	Alias       *string
	ExpiresAt   *time.Time
	MaxClicks   *int
}

func (uc *LinkUC) Create(ctx context.Context, req CreateLinkRequest) (*domain.Link, error) {
	if req.OriginalURL == "" {
		return nil, errors.New("original_url required")
	}
	alias := ""
	if req.Alias != nil && *req.Alias != "" {
		if !uc.alias.ValidCustom(*req.Alias) {
			return nil, errors.New("invalid alias format")
		}
		alias = *req.Alias
	} else {
		alias = uc.alias.Generate(7)
	}

	l := &domain.Link{
		ID:          uuid.NewString(),
		UserID:      req.UserID,
		OriginalURL: req.OriginalURL,
		Alias:       alias,
		ExpiresAt:   req.ExpiresAt,
		MaxClicks:   req.MaxClicks,
		IsActive:    true,
		ClickCount:  0,
		CreatedAt:   uc.now(),
		UpdatedAt:   uc.now(),
	}

	if err := uc.repo.Create(ctx, l); err != nil {
		return nil, err
	}
	return l, nil
}
