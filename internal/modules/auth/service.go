package auth

import (
	"context"
	"errors"
	"fmt"
	"restaurant-management/internal/modules/user"
	"restaurant-management/pkg/jwt"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	jwt      *jwtauth.JWTAuth
	userRepo user.Repository
}

var (
	ErrUserAlreadyExist = errors.New("user already exists")
	ErrUserNotFound     = errors.New("user not found")
)

func NewService(userRepo user.Repository, jwt *jwtauth.JWTAuth) *Service {
	return &Service{userRepo: userRepo, jwt: jwt}
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	// Cek Email
	u, err := s.userRepo.GetByEmail(ctx, req.Email)

	if err != nil {
		if !errors.Is(err, user.ErrUserNotFound) {
			return LoginResponse{}, ErrUserNotFound
		}
		return LoginResponse{}, fmt.Errorf("failed to get user by email: %w", err)
	}

	// Cek Password
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return LoginResponse{}, fmt.Errorf("invalild password: %w", err)
	}

	// Create Token String
	claims := jwt.CustomClaims{
		UserID: u.ID,
		Email:  u.Email,
		Role:   "user",
	}
	_, tokenString, err := claims.GenerateToken(s.jwt, time.Now().Add(24*time.Hour))
	if err != nil {
		return LoginResponse{}, err
	}

	// Create Refresh Token
	_, refreshTokenString, err := claims.GenerateToken(s.jwt, time.Now().Add(7*24*time.Hour))
	if err != nil {
		return LoginResponse{}, err
	}

	u.Password = ""

	// Return Login Repsonse
	return LoginResponse{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
		User:         &u,
	}, nil
}

func (s *Service) Register(ctx context.Context, req RegisterRequest) error {
	// Cek Email
	_, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err == nil {
		return ErrUserAlreadyExist
	}

	if !errors.Is(err, user.ErrUserNotFound) {
		return fmt.Errorf("failed to check email: %w", err)
	}

	// Hash Passworod
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password")
	}

	u := user.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  string(hashPassword),
		Phone:     req.Phone,
		Avatar:    req.Avatar,
	}
	if _, err := s.userRepo.Create(ctx, u); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *Service) RefreshToken(ctx context.Context, refreshToken string) (LoginResponse, error) {
	c, err := jwt.GetClaims(ctx)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("invalid refresh token", err)
	}

	u, err := s.userRepo.GetByID(ctx, int(c.UserID))
	if err != nil {
		return LoginResponse{}, err
	}

	claims := jwt.CustomClaims{
		UserID: u.ID,
		Email:  u.Email,
		Role:   "user",
	}

	_, tokenString, err := claims.GenerateToken(s.jwt, time.Now().Add(24*time.Hour))
	if err != nil {
		return LoginResponse{}, err
	}

	// Create Refresh Token
	_, refreshTokenString, err := claims.GenerateToken(s.jwt, time.Now().Add(7*24*time.Hour))
	if err != nil {
		return LoginResponse{}, err
	}

	// Return Login Repsonse
	return LoginResponse{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
	}, nil

}
