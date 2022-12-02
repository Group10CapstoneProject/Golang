package myerrors

import "errors"

var (
	ErrEmailAlredyExist     = errors.New("email is used")
	ErrUserNotFound         = errors.New("user not found")
	ErrInvalidEmailPassword = errors.New("invalid email or password")
	ErrGenerateAccessToken  = errors.New("error when generate access token")
	ErrGenerateRefreshToken = errors.New("error when generate refresh token")
	ErrToken                = errors.New("invalid token")
	ErrTokenExpired         = errors.New("token expired")
	ErrPermission           = errors.New("not have permission to access")
	ErrInvalidSession       = errors.New("invalid session id")
	ErrDuplicateRecord      = errors.New("duplicate record")
	ErrRecordNotFound       = errors.New("record not found")
	ErrFailedUpload         = errors.New("upload file failed")
)
