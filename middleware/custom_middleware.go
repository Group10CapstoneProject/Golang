package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Group10CapstoneProject/Golang/constans"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type customMiddleware struct {
	db     *gorm.DB
	secret string
}

func (j *customMiddleware) CustomJWTWithConfig(role string) echo.MiddlewareFunc {
	config := middleware.DefaultJWTConfig
	config.SigningKey = []byte(j.secret)
	config.KeyFunc = func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if t.Method.Alg() != config.SigningMethod {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		if len(config.SigningKeys) > 0 {
			if kid, ok := t.Header["kid"].(string); ok {
				if key, ok := config.SigningKeys[kid]; ok {
					return key, nil
				}
			}
			return nil, fmt.Errorf("unexpected jwt key id=%v", t.Header["kid"])
		}

		return config.SigningKey, nil
	}
	config.ParseTokenFunc = func(auth string, c echo.Context) (interface{}, error) {
		token, err := jwt.Parse(auth, config.KeyFunc)
		if err != nil {
			fmt.Println("sukses")
			return nil, err
		}
		if !token.Valid {
			return nil, errors.New("invalid token")
		}
		claims := token.Claims.(jwt.MapClaims)
		if role != "" {
			userRole := claims["role"].(string)
			if userRole != role && userRole != constans.Role_superadmin {
				return nil, myerrors.ErrPermission
			}
		}
		userId := uint(claims["user_id"].(float64))
		user := model.User{ID: userId}
		if err := j.db.First(&user).Error; err != nil {
			return nil, err
		}
		sessionId := claims["session_id"].(string)
		if user.SessionID.String() != sessionId {
			return nil, myerrors.ErrInvalidSession
		}
		return token, nil
	}
	config.ErrorHandler = func(err error) error {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		if err == myerrors.ErrInvalidSession {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return middleware.JWTWithConfig(config)
}

func NewCustomMiddleware(db *gorm.DB, secret string) customMiddleware {
	return customMiddleware{
		db:     db,
		secret: secret,
	}
}
