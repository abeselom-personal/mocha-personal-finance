package models

import (
	"time"
)

type Token struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index;not null"`
	Token     string `gorm:"uniqueIndex;not null"`
	Type      string `gorm:"not null"` // access, refresh
	Scope     string // optional scope/claim
	IP        string
	UserAgent string
	ExpiresAt time.Time
	Revoked   bool `gorm:"default:false"`
	CreatedAt time.Time
}
