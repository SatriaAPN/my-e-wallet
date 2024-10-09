package service

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

type CreateUserResponse struct {
	Name  string
	Email string
}

type LoginUserRequest struct {
	Email    string
	Password string
}

type LoginUserResponse struct {
	Token string
}

type ProfileUserResponse struct {
	Name          string
	Email         string
	WalletNumber  string
	WalletBalance decimal.Decimal
}

type ForgetPasswordRequest struct {
	Email string
}

type ForgetPasswordResponse struct {
	Token     string
	ExpiredAt time.Time
}

type ResetPasswordRequest struct {
	Email       string
	NewPassword string
	Token       string
}
