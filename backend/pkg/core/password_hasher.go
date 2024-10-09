package core

import (
	"github.com/SatriaAPN/my-e-wallet/backend/pkg/config"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	GenerateHashFromPassword(password string) (string, error)
	CompareHashAndPassword(password string, hashedPassword string) (bool, error) // return true if password is matched to the hashed password
}

func GetPasswordHasher() PasswordHasher {
	if passwordHasherInstance == nil {
		passwordHasherInstance = newPasswordHasher()
	}

	return passwordHasherInstance
}

func newPasswordHasher() *passwordHasher {
	return &passwordHasher{}
}

type passwordHasher struct{}

var passwordHasherInstance *passwordHasher

func (ph *passwordHasher) GenerateHashFromPassword(password string) (string, error) {
	result := ""

	pByte, err := bcrypt.GenerateFromPassword([]byte(password), config.BcryptCost())

	result = string(pByte)

	if err != nil {
		return result, err
	}

	return result, nil
}

// return true if the password match the hashed pasword
func (ph *passwordHasher) CompareHashAndPassword(password string, hashedPassword string) (bool, error) {
	result := true

	pByte := []byte(password)
	hpByte := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(hpByte, pByte)

	if err != nil {
		return false, err
	}

	return result, nil
}
