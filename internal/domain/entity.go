package domain

import "time"

type User struct {
	ID           string    `gorm:"type:uuid;primaryKey" json:"id"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type Link struct {
	ID           string     `gorm:"type:uuid;primaryKey" json:"id"`
	UserID       string     `gorm:"type:uuid;index;not null" json:"user_id"`
	OriginalURL  string     `gorm:"not null" json:"original_url"`
	Alias        string     `gorm:"size:64;uniquieIndex;not null" json:"alias"`
	PasswordHash *string    `json:"-"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
	MaxClicks    *int       `json:"max_clicks,omitempty"`
	IsActive     bool       `gorm:"default:true" json:"is_active"`
	ClickCount   int64      `gorm:"default:0" json:"click_count"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type Click struct {
	ID       int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	LinkID   string    `gorm:"type:uuid;index;not null" json:"link_id"`
	TS       time.Time `gorm:"autoCreateTime" json:"ts"`
	IPHash   string    `gorm:"index;not null" json:"ip_hash"`
	Country  *string   `json:"country,omitempty"`
	City     *string   `json:"city,omitempty"`
	Referrer *string   `json:"referrer,omitempty"`
	UA       *string   `json:"ua,omitempty"`
	Device   *string   `json:"device,omitempty"`
	OS       *string   `json:"os,omitempty"`
	Browser  *string   `json:"browser,omitempty"`
}
