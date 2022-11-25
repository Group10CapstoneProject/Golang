package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/auth/dto"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthService interface {
	Login(credential dto.UserCredential, ctx context.Context) (*model.Token, error)
	LoginAdmin(credential dto.UserCredential, ctx context.Context) (*model.Token, error)
	Logout(userID uint, ctx context.Context) error
	Refresh(token model.Token, ctx context.Context) (*model.Token, error)
	ValidationToken(token jwt.MapClaims, ctx context.Context) error
	ValidationRole(token jwt.MapClaims, role string, ctx context.Context) error
	GetClaims(c *echo.Context) jwt.MapClaims
}
