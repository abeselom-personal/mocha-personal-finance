package models

type Role struct {
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"uniqueIndex;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}
