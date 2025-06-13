package models

type Permission struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"uniqueIndex;not null"` // e.g. "user:create"
	Group       string `gorm:"index;not null"`       // e.g. "user", "auth"
	Description string
	Roles       []Role `gorm:"many2many:role_permissions"`
	Users       []User `gorm:"many2many:user_permissions"`
}
