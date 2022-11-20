package jwt

import (
	"time"

	"github.com/Group10CapstoneProject/Golang/constans"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	generateAccessToken(user *model.User) (string, error)
	generateRefreshToken(user *model.User) (string, error)
	GenerateToken(user *model.User) (at string, rt string, err error)
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

func (j *jwtServiceImpl) generateAccessToken(user *model.User) (string, error) {
	claims := &jwt.MapClaims{
		"user_id":    user.ID,
		"role":       user.Role,
		"session_id": user.SessionID,
		"exp":        time.Now().Add(constans.ExpAccessToken).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.accessSecretKey))
}

func (j *jwtServiceImpl) generateRefreshToken(user *model.User) (string, error) {
	claims := &jwt.MapClaims{
		"user_id":    user.ID,
		"session_id": user.SessionID,
		"exp":        time.Now().Add(constans.ExpRefreshToken).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.refreshSecretKey))
}

func (j *jwtServiceImpl) GenerateToken(user *model.User) (at string, rt string, err error) {
	accessToken, err := j.generateAccessToken(user)
	if err != nil {
		return "", "", myerrors.ErrGenerateAccessToken
	}
	refreshToken, err := j.generateRefreshToken(user)
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
