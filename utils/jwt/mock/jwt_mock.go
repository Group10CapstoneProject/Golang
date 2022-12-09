package mock

import (
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockJWTService struct {
	mock.Mock
}

func (m *MockJWTService) GenerateToken(user *model.User) (string, string, error) {
	args := m.Called()
	return args.String(0), args.String(1), args.Error(2)
}

func (m *MockJWTService) GetClaims(c *echo.Context) jwt.MapClaims {
	args := m.Called()

	return args.Get(0).(jwt.MapClaims)
}

func (m *MockJWTService) TokenClaims(token string, secret string) (jwt.MapClaims, error) {
	args := m.Called()

	return args.Get(0).(jwt.MapClaims), args.Error(1)
}

func (m *MockJWTService) GenerateAccessToken(user *model.User) (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockJWTService) GenerateRefreshToken(user *model.User) (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}
