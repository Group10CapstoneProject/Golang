package mock

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/internal/auth/dto"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type AuthServiceMock struct {
	mock.Mock
}

func (m *AuthServiceMock) Login(credential dto.UserCredential, ctx context.Context) (*model.Token, error) {
	args := m.Called()
	return args.Get(0).(*model.Token), args.Error(1)
}
func (m *AuthServiceMock) LoginAdmin(credential dto.UserCredential, ctx context.Context) (*model.AdminToken, error) {
	args := m.Called()
	return args.Get(0).(*model.AdminToken), args.Error(1)
}
func (m *AuthServiceMock) Logout(userID uint, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}
func (m *AuthServiceMock) Refresh(token model.Token, ctx context.Context) (*model.Token, error) {
	args := m.Called()
	return args.Get(0).(*model.Token), args.Error(1)
}
func (m *AuthServiceMock) ValidationToken(token jwt.MapClaims, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}
func (m *AuthServiceMock) ValidationRole(token jwt.MapClaims, role string, ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}
func (m *AuthServiceMock) GetClaims(c *echo.Context) jwt.MapClaims {
	args := m.Called()
	return args.Get(0).(jwt.MapClaims)
}
