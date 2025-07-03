// service/auth_service.go
package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/abeselom-personal/personal-finance/config"
	"github.com/abeselom-personal/personal-finance/dto"
	"github.com/abeselom-personal/personal-finance/models"
	"github.com/abeselom-personal/personal-finance/repositories"
	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrEmailExists     = errors.New("email already exists")
	ErrUsernameExists  = errors.New("username already exists")
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidToken    = errors.New("invalid token")
	ErrTokenRevoked    = errors.New("token revoked")
	ErrTokenExpired    = errors.New("token expired")
)

type AuthService struct {
	UserRepo  repositories.UserRepository
	TokenRepo repositories.TokenRepository
	RoleRepo  repositories.RoleRepository
	Config    *config.Config
}

func NewAuthService(
	userRepo repositories.UserRepository,
	tokenRepo repositories.TokenRepository,
	roleRepo repositories.RoleRepository,
	cfg *config.Config,
) *AuthService {
	return &AuthService{
		UserRepo:  userRepo,
		TokenRepo: tokenRepo,
		RoleRepo:  roleRepo,
		Config:    cfg,
	}
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) (*dto.TokenResponse, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	existingEmail, err := s.UserRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email: %w", err)
	}
	if existingEmail != nil {
		return nil, fmt.Errorf("%w: %s", ErrEmailExists, email)
	}

	existingUsername, err := s.UserRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to check username: %w", err)
	}
	if existingUsername != nil {
		return nil, fmt.Errorf("%w: %s", ErrUsernameExists, req.Username)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), s.Config.BcryptCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	role, err := s.RoleRepo.GetByName(ctx, "user")
	if err != nil {
		return nil, fmt.Errorf("failed to get default role: %w", err)
	}
	if role == nil {
		return nil, errors.New("default role not found")
	}

	user := &models.User{
		Username:     req.Username,
		Email:        email,
		PasswordHash: string(hashed),
		FullName:     req.FullName,
		RoleID:       role.ID,
		IsActive:     true,
	}

	if err := s.UserRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return s.generateTokens(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (*dto.TokenResponse, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	user, err := s.UserRepo.GetByEmail(ctx, email)
	if err != nil || user == nil {
		return nil, fmt.Errorf("%w: user not found", ErrInvalidPassword)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPassword, err)
	}

	now := time.Now()
	user.LastLogin = &now
	if err := s.UserRepo.Update(ctx, user); err != nil {
		fmt.Printf("failed to update last login: %v\n", err)
	}

	return s.generateTokens(ctx, user)
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*dto.TokenResponse, error) {

	tokenRecord, err := s.TokenRepo.GetByToken(ctx, refreshToken)
	fmt.Println(refreshToken)
	fmt.Println(err)
	if err != nil || tokenRecord == nil || tokenRecord.Revoked {
		return nil, ErrTokenRevoked
	}

	user, err := s.UserRepo.GetByID(ctx, tokenRecord.UserID)
	if err != nil || user == nil {
		return nil, fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}

	accessToken, accessExp, err := s.generateAccessToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	return &dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int64(time.Until(accessExp).Seconds()),
	}, nil
}

func (s *AuthService) RevokeToken(ctx context.Context, token string) error {
	if err := s.TokenRepo.Revoke(ctx, token); err != nil {
		return fmt.Errorf("failed to revoke token: %w", err)
	}
	return nil
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) (*Claims, error) {
	claims, err := s.validateToken(token)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	if claims.TokenType != "access" {
		return nil, fmt.Errorf("%w: not an access token", ErrInvalidToken)
	}

	tokenRecord, err := s.TokenRepo.GetByToken(ctx, token)
	if err == nil && tokenRecord != nil && tokenRecord.Revoked {
		return nil, ErrTokenRevoked
	}

	return claims, nil
}

func (s *AuthService) generateTokens(ctx context.Context, user *models.User) (*dto.TokenResponse, error) {
	accessToken, accessExp, err := s.generateAccessToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, refreshExp, err := s.generateRefreshToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	if err := s.storeRefreshToken(ctx, user.ID, refreshToken, refreshExp); err != nil {
		return nil, fmt.Errorf("failed to store refresh token: %w", err)
	}

	return &dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int64(time.Until(accessExp).Seconds()),
	}, nil
}

func (s *AuthService) generateAccessToken(user *models.User) (string, time.Time, error) {
	expiration := time.Now().Add(s.Config.AccessTokenExpiration)
	claims := &Claims{
		UserID:    strconv.FormatUint(uint64(user.ID), 10),
		Username:  user.Username,
		Role:      user.Role.Name,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.Config.JWTSecret))
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to sign access token: %w", err)
	}

	return signedToken, expiration, nil
}

func (s *AuthService) generateRefreshToken(_ uint) (string, time.Time, error) {
	expiration := time.Now().Add(s.Config.RefreshTokenExpiration)
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", time.Time{}, fmt.Errorf("failed to generate random bytes: %w", err)
	}
	token := base64.URLEncoding.EncodeToString(tokenBytes)
	return token, expiration, nil
}

func (s *AuthService) storeRefreshToken(ctx context.Context, userID uint, token string, expiresAt time.Time) error {
	refreshToken := &models.Token{
		UserID:    userID,
		Token:     token,
		Type:      "refresh",
		ExpiresAt: expiresAt,
		Revoked:   false,
	}
	if err := s.TokenRepo.Create(ctx, refreshToken); err != nil {
		return fmt.Errorf("failed to store refresh token: %w", err)
	}
	return nil
}

func (s *AuthService) validateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.Config.JWTSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("token parsing failed: %w", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

type Claims struct {
	UserID    string `json:"user_id"` // ← change to string
	Username  string `json:"username"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}
