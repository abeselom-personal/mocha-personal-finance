package service

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/abeselom-personal/personal-finance/config"
	"github.com/abeselom-personal/personal-finance/dto"
	"github.com/abeselom-personal/personal-finance/models"
	"github.com/abeselom-personal/personal-finance/repositories"
	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	UserRepo  *repositories.UserRepository
	RoleRepo  *repositories.RoleRepository
	TokenRepo *repositories.TokenRepository
	JWTSecret string
}

func NewAuthService(userRepo *repositories.UserRepository, roleRepo *repositories.RoleRepository, tokenRepo *repositories.TokenRepository) *AuthService {
	return &AuthService{
		UserRepo:  userRepo,
		RoleRepo:  roleRepo,
		TokenRepo: tokenRepo,
		JWTSecret: config.Cfg.JWTString,
	}
}

func (s *AuthService) Register(req dto.RegisterDTO) (string, error) {
	existing, _ := s.UserRepo.GetByUsername(req.Username)
	if existing != nil {
		return "", errors.New("username already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return "", err
	}

	role, err := s.RoleRepo.GetByName("user")
	if err != nil {
		return "", errors.New("default role not found")
	}

	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashed),
		FullName:     req.FullName,
		RoleID:       role.ID,
		IsActive:     true,
	}

	if err := s.UserRepo.Create(user); err != nil {
		return "", err
	}

	return s.generateJWT(user)
}

func (s *AuthService) Login(req dto.LoginDTO) (string, error) {
	user, err := s.UserRepo.GetByUsername(req.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	now := time.Now()
	user.LastLogin = &now
	_ = s.UserRepo.Update(user)

	return s.generateJWT(user)
}

func (s *AuthService) generateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role.Name,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.JWTSecret))
}
