package jwt

import (
	"time"

	"github.com/Group10CapstoneProject/Golang/constans"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JWTService interface {
	GenerateAccessToken(user *model.User) (string, error)
	GenerateRefreshToken(user *model.User) (string, error)
	GenerateToken(user *model.User) (at string, rt string, err error)
	GetClaims(c *echo.Context) jwt.MapClaims
	TokenClaims(token string, secret string) (jwt.MapClaims, error)
}

type jwtServiceImpl struct {
	accessSecretKey  string
	refreshSecretKey string
}

func NewJWTService(accessSecretKey string, refreshSecretKey string) JWTService {
	return &jwtServiceImpl{
		accessSecretKey:  accessSecretKey,
		refreshSecretKey: refreshSecretKey,
	}
}

func (j *jwtServiceImpl) GenerateAccessToken(user *model.User) (string, error) {
	claims := &jwt.MapClaims{
		"user_id":    user.ID,
		"role":       user.Role,
		"session_id": user.SessionID,
		"exp":        time.Now().Add(constans.ExpAccessToken).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.accessSecretKey))
}

func (j *jwtServiceImpl) GenerateRefreshToken(user *model.User) (string, error) {
	claims := &jwt.MapClaims{
		"user_id":    user.ID,
		"session_id": user.SessionID,
		"exp":        time.Now().Add(constans.ExpRefreshToken).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.refreshSecretKey))
}

func (j *jwtServiceImpl) GenerateToken(user *model.User) (at string, rt string, err error) {
	accessToken, err := j.GenerateAccessToken(user)
	if err != nil {
		return "", "", myerrors.ErrGenerateAccessToken
	}
	refreshToken, err := j.GenerateRefreshToken(user)
	if err != nil {
		return "", "", myerrors.ErrGenerateRefreshToken
	}
	return accessToken, refreshToken, nil
}

func (j *jwtServiceImpl) TokenClaims(token string, secret string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (j *jwtServiceImpl) GetClaims(c *echo.Context) jwt.MapClaims {
	user := (*c).Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}
