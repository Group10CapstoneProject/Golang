package jwt

import (
	"time"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JWTService interface {
	GenerateToken(user *model.User) (string, error)
	GetClaims(c *echo.Context) jwt.MapClaims
}

type jwtServiceImpl struct {
	secretKey string
	exp       time.Duration
}

func NewJWTService(secretKey string, exp time.Duration) JWTService {
	return &jwtServiceImpl{
		secretKey: secretKey,
		exp:       exp,
	}
}

func (j *jwtServiceImpl) GenerateToken(user *model.User) (string, error) {
	claims := &jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(j.exp).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtServiceImpl) GetClaims(c *echo.Context) jwt.MapClaims {
	user := (*c).Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}
