// repositories/user_repository.go
package repositories

import (
	"context"

	"github.com/abeselom-personal/personal-finance/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Preload("Role").
		First(&user, id).Error
	return &user, err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		Preload("Role").
		First(&user).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Where("username = ?", username).
		Preload("Role").
		First(&user).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, id).Error
}

func (r *UserRepository) AddPermission(ctx context.Context, userID uint, permissionID uint) error {
	return r.db.WithContext(ctx).Exec(
		"INSERT INTO user_permissions (user_id, permission_id) VALUES (?, ?)",
		userID,
		permissionID,
	).Error
}

func (r *UserRepository) RemovePermission(ctx context.Context, userID uint, permissionID uint) error {
	return r.db.WithContext(ctx).Exec(
		"DELETE FROM user_permissions WHERE user_id = ? AND permission_id = ?",
		userID,
		permissionID,
	).Error
}
