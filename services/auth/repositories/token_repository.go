// repositories/token_repository.go
package repositories

import (
	"context"

	"github.com/abeselom-personal/personal-finance/models"
	"gorm.io/gorm"
)

type TokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (r *TokenRepository) Create(ctx context.Context, token *models.Token) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *TokenRepository) GetByToken(ctx context.Context, token string) (*models.Token, error) {
	var t models.Token
	err := r.db.WithContext(ctx).
		Where("token = ?", token).
		First(&t).Error
	return &t, err
}

func (r *TokenRepository) Revoke(ctx context.Context, token string) error {
	result := r.db.WithContext(ctx).
		Model(&models.Token{}).
		Where("token = ?", token).
		Update("revoked", true)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *TokenRepository) RevokeAllForUser(ctx context.Context, userID uint) error {
	return r.db.WithContext(ctx).
		Model(&models.Token{}).
		Where("user_id = ?", userID).
		Update("revoked", true).Error
}
