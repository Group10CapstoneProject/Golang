package password

import "golang.org/x/crypto/bcrypt"

type PasswordHash interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type password struct{}

func (*password) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (*password) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewPasswordService() PasswordHash {
	return &password{}
}
