package core

import "net/mail"

type DataValidator interface {
	IsEmailValid(email string) bool
}

type dataValidator struct{}

func NewDataValidator() DataValidator {
	return &dataValidator{}
}

func (dv *dataValidator) IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
