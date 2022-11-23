package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/config"
	"github.com/Group10CapstoneProject/Golang/constans"
	"github.com/Group10CapstoneProject/Golang/internal/auth/dto"
	authRepository "github.com/Group10CapstoneProject/Golang/internal/auth/repository"
	usersRepository "github.com/Group10CapstoneProject/Golang/internal/users/repository"
	"github.com/Group10CapstoneProject/Golang/model"
	jwtService "github.com/Group10CapstoneProject/Golang/utils/jwt"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/Group10CapstoneProject/Golang/utils/password"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type authServiceImpl struct {
	authRepository  authRepository.AuthRepository
	usersRepository usersRepository.UserRepository
	password        password.PasswordHash
	jwtService      jwtService.JWTService
}

// Login implements AuthService
func (s *authServiceImpl) Login(credential dto.UserCredential, ctx context.Context) (*model.Token, error) {
	user, err := s.usersRepository.FindUserByEmail(&credential.Email, ctx)
	if err != nil {
		return nil, myerrors.ErrInvalidEmailPassword
	}
	if !s.password.CheckPasswordHash(credential.Password, user.Password) {
		return nil, myerrors.ErrInvalidEmailPassword
	}
	user.SessionID = uuid.New()
	var newToken model.Token
	newToken.AccessToken, newToken.RefreshToken, err = s.jwtService.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	if err := s.authRepository.UpdateSessionID(user.ID, user.SessionID, ctx); err != nil {
		return nil, err
	}
	return &newToken, err
}

// Logout implements AuthService
func (s *authServiceImpl) Logout(userID uint, ctx context.Context) error {
	err := s.authRepository.UpdateSessionID(userID, uuid.Nil, ctx)
	return err
}

// Refresh implements AuthService
func (s *authServiceImpl) Refresh(token model.Token, ctx context.Context) (*model.Token, error) {
	refreshToken, err := s.jwtService.TokenClaims(token.RefreshToken, config.Env.JWT_SECRET_REFRESH)
	if err != nil {
		return nil, err
	}
	// compare session id
	userId := uint(refreshToken["user_id"].(float64))
	sessionId := refreshToken["session_id"].(string)
	user, err := s.usersRepository.FindUserByID(&userId, ctx)
	if err != nil {
		return nil, err
	}
	if sessionId != user.SessionID.String() {
		return nil, myerrors.ErrToken
	}
	// generate new token
	user.SessionID = uuid.New()
	var newToken model.Token
	newToken.AccessToken, newToken.RefreshToken, err = s.jwtService.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	if err := s.authRepository.UpdateSessionID(user.ID, user.SessionID, ctx); err != nil {
		return nil, err
	}
	return &newToken, err
}

// ValidatationAdmin implements AuthService
func (s *authServiceImpl) ValidationRole(token jwt.MapClaims, role string, ctx context.Context) error {
	userRole := token["role"].(string)
	if userRole == role || userRole == constans.Role_superadmin {
		return s.ValidationToken(token, ctx)
	}
	return myerrors.ErrPermission
}

// ValidationToken implements AuthService
func (s *authServiceImpl) ValidationToken(token jwt.MapClaims, ctx context.Context) error {
	userId := uint(token["user_id"].(float64))
	user, err := s.usersRepository.FindUserByID(&userId, ctx)
	if err != nil {
		return err
	}
	sessionId := token["session_id"].(string)
	if user.SessionID.String() != sessionId {
		return myerrors.ErrInvalidSession
	}
	return nil
}

func (s *authServiceImpl) GetClaims(c *echo.Context) jwt.MapClaims {
	user := (*c).Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}

func NewAuthService(auth authRepository.AuthRepository, users usersRepository.UserRepository, password password.PasswordHash, jwt jwtService.JWTService) AuthService {
	return &authServiceImpl{
		authRepository:  auth,
		usersRepository: users,
		password:        password,
		jwtService:      jwt,
	}
}
