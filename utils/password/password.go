package password

import "golang.org/x/crypto/bcrypt"

type PasswordHash interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type Password struct{}

func (Password) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (Password) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
