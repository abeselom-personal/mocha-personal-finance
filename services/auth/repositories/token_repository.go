package repositories

import (
	"errors"

	"github.com/abeselom-personal/personal-finance/models"
	"gorm.io/gorm"
)

type TokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (r *TokenRepository) Create(token *models.Token) error {
	return r.db.Create(token).Error
}

func (r *TokenRepository) GetByToken(token string) (*models.Token, error) {
	var t models.Token
	err := r.db.Where("token = ?", token).First(&t).Error
	return &t, err
}

func (r *TokenRepository) Revoke(token string) error {
	result := r.db.Model(&models.Token{}).
		Where("token = ?", token).
		Update("revoked", true)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("token not found")
	}
	return nil
}

func (r *TokenRepository) RevokeAllForUser(userID uint) error {
	return r.db.Model(&models.Token{}).
		Where("user_id = ?", userID).
		Update("revoked", true).Error
}
