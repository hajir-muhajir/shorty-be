package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/hajir.muhajir/shorty-api/internal/domain"
	"github.com/hajir.muhajir/shorty-api/internal/repository"
)

type RedirectUC struct {
	links repository.LinkRepository
	click repository.ClickRepository
	now   func() time.Time
}

func NewRedirectUC(l repository.LinkRepository, c repository.ClickRepository) *RedirectUC {
	return &RedirectUC{
		links: l,
		click: c,
		now:   time.Now,
	}
}

func (uc *RedirectUC) Resolve(ctx context.Context, alias string) (*domain.Link, error) {
	link, err := uc.links.FindByAlias(ctx, alias)
	if err != nil {
		return nil, err
	}

	if !link.IsActive {
		return nil, domain.ErrInactive
	}
	if link.ExpiresAt != nil && uc.now().After(*link.ExpiresAt) {
		return nil, domain.ErrExpired
	}
	if link.MaxClicks != nil && int64(*link.MaxClicks) <= link.ClickCount {
		return nil, domain.ErrLimitReached
	}
	return link, nil
}

func (uc *RedirectUC) LogClick(ctx context.Context, link *domain.Link, ip, referrer, ua string) error {
	sum := sha256.Sum256([]byte(ip))
	iphash := hex.EncodeToString(sum[:])

	c := &domain.Click{
		LinkID:   link.ID,
		TS:       uc.now(),
		IPHash:   iphash,
		Referrer: strPtrOrNil(referrer),
		UA:       strPtrOrNil(ua),
	}

	if err := uc.click.Insert(ctx, c); err != nil {
		return err
	}
	if err := uc.links.IncrementClick(ctx, link.ID); err != nil {
		return err
	}
	return uc.links.SetUpdateAt(ctx, link.ID, uc.now())

}

func strPtrOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
