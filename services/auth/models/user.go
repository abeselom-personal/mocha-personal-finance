package models

import (
	"time"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"uniqueIndex;size:255;not null"`
	Email        string `gorm:"uniqueIndex;size:255;not null"`
	PasswordHash string `gorm:"type:varchar(255);not null"`
	FullName     string `gorm:"size:255"`
	RoleID       uint
	Role         Role
	Permissions  []Permission `gorm:"many2many:user_permissions;constraint:OnDelete:CASCADE"`
	Tokens       []Token      `gorm:"constraint:OnDelete:CASCADE"`
	IsActive     bool         `gorm:"default:true"`
	LastLogin    *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `swaggerignore:"true"`
}
