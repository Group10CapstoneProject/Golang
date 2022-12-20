package myerrors

import (
	"errors"
	"regexp"
	"strings"
)

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
	ErrAlredyPaid           = errors.New("alredy paid")
	ErrAlredyTake           = errors.New("booking alredy take")
	ErrRecordIsUsed         = errors.New("record is used by other tables")
	ErrPaymentMethod        = errors.New("payment method declined")
	ErrOrderExpired         = errors.New("order is expired")
	ErrIsCanceled           = errors.New("order is canceled")
	ErrCantCanceled         = errors.New("order can't canceled")
	ErrTrainerIsFull        = errors.New("trainer slot is full for this day")
	ErrForeignKey           = func(err error) error {
		src := regexp.MustCompile(`REFERENCES`)
		temp := src.Split(err.Error(), -1)
		massage := strings.Replace(temp[1], "(", "", -1)
		massage = strings.Replace(massage, ")", "", -1)
		return errors.New("invalid" + massage)
	}
)
