package repositories

import (
	"context"

	"github.com/abeselom-personal/personal-finance/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) Create(ctx context.Context, role *models.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *RoleRepository) GetByName(ctx context.Context, name string) (*models.Role, error) {
	var role models.Role
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&role).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &role, err
}

func (r *RoleRepository) AddPermission(ctx context.Context, roleID uint, permissionID uint) error {
	return r.db.WithContext(ctx).Exec(
		"INSERT INTO role_permissions (role_id, permission_id) VALUES (?, ?)",
		roleID,
		permissionID,
	).Error
}
