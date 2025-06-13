package repositories

import (
	"github.com/abeselom-personal/personal-finance/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").Preload("Permissions").First(&user, id).Error
	return &user, err
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).
		Preload("Role").
		Preload("Permissions").
		First(&user).Error
	return &user, err
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *UserRepository) AddPermission(userID uint, permissionID uint) error {
	return r.db.Exec(
		"INSERT INTO user_permissions (user_id, permission_id) VALUES (?, ?)",
		userID,
		permissionID,
	).Error
}

func (r *UserRepository) RemovePermission(userID uint, permissionID uint) error {
	return r.db.Exec(
		"DELETE FROM user_permissions WHERE user_id = ? AND permission_id = ?",
		userID,
		permissionID,
	).Error
}

func (r *UserRepository) WithTx(tx *gorm.DB) *UserRepository {
	return &UserRepository{db: tx}
}
