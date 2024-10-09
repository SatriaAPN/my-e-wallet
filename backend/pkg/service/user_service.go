package service

import (
	"context"

	"github.com/SatriaAPN/my-e-wallet/backend/pkg/core/constants"
	coreerrors "github.com/SatriaAPN/my-e-wallet/backend/pkg/core/errors"
	"github.com/SatriaAPN/my-e-wallet/backend/pkg/repository"

	"github.com/go-errors/errors"

	core "github.com/SatriaAPN/my-e-wallet/backend/pkg/core"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, cu CreateUserRequest) (CreateUserResponse, error)
	LoginUser(ctx context.Context, cu LoginUserRequest) (LoginUserResponse, error)
}

type UserUsecaseConfig struct {
	UserRepository       repository.UserRepository
	PasswordHasher       core.PasswordHasher
	AuthTokenGenerator   core.AuthTokenGenerator
	RandomTokenGenerator core.RandomTokenGenerator
	DataValidator        core.DataValidator
}

func NewUserUsecase(config UserUsecaseConfig) UserUsecase {
	uu := userUsecase{}

	if config.UserRepository != nil {
		uu.userRepository = config.UserRepository
	}
	if config.PasswordHasher != nil {
		uu.passwordHasher = config.PasswordHasher
	}
	if config.AuthTokenGenerator != nil {
		uu.authTokenGenerator = config.AuthTokenGenerator
	}
	if config.RandomTokenGenerator != nil {
		uu.randomTokenGenerator = config.RandomTokenGenerator
	}

	return &uu
}

type userUsecase struct {
	userRepository repository.UserRepository

	authTokenGenerator   core.AuthTokenGenerator
	passwordHasher       core.PasswordHasher
	randomTokenGenerator core.RandomTokenGenerator
	dataValidator        core.DataValidator
}

func (uu *userUsecase) CreateUser(ctx context.Context, cu CreateUserRequest) (CreateUserResponse, error) {
	cures := CreateUserResponse{}

	if err := uu.checkCreateUserData(ctx, cu); err != nil {
		return cures, coreerrors.ErrorHandling(err)
	}

	hashedPassword, err := uu.passwordHasher.GenerateHashFromPassword(cu.Password)

	if err != nil {
		return cures, coreerrors.ErrorHandling(err)
	}

	u := core.User{
		Name:     cu.Name,
		Email:    cu.Email,
		Password: hashedPassword,
	}

	u2, err := uu.userRepository.Create(ctx, u)

	if err != nil {
		return cures, coreerrors.ErrorHandling(err)
	}

	cures.Name = u2.Name
	cures.Email = u2.Email

	return cures, nil
}

func (uu *userUsecase) checkCreateUserData(ctx context.Context, cu CreateUserRequest) error {
	if !uu.dataValidator.IsEmailValid(cu.Email) {
		return errors.New(coreerrors.ErrEmailIsNotValid)
	}

	if len(cu.Password) < constants.MinimumPasswordLength {
		return errors.New(coreerrors.ErrMinimumPasswordLength)
	}

	if len(cu.Password) > constants.MaximumPasswordLength {
		return errors.New(coreerrors.ErrMaximumPasswordLength)
	}

	return nil
}

func (uu *userUsecase) LoginUser(ctx context.Context, cu LoginUserRequest) (LoginUserResponse, error) {
	lures := LoginUserResponse{}

	if err := uu.checkLoginUserData(ctx, cu); err != nil {
		return lures, coreerrors.ErrorHandling(err)
	}

	u, err := uu.userRepository.FindByEmail(ctx, cu.Email)

	if err != nil {
		return lures, coreerrors.ErrorHandling(err)
	}

	if u.ID == 0 {
		return lures, errors.New(coreerrors.ErrEmailNotFound)
	}

	match, err := uu.passwordHasher.CompareHashAndPassword(cu.Password, u.Password)

	if !match {
		return lures, errors.New(coreerrors.ErrWrongPassword)
	}

	if err != nil {
		return lures, coreerrors.ErrorHandling(err)
	}

	ad := core.AuthData{
		ID: u.ID,
	}

	token, err := uu.authTokenGenerator.Encode(ad)

	if err != nil {
		return lures, coreerrors.ErrorHandling(err)
	}

	lures.Token = token

	return lures, nil
}

func (uu *userUsecase) checkLoginUserData(ctx context.Context, cu LoginUserRequest) error {
	if !uu.dataValidator.IsEmailValid(cu.Email) {
		return errors.New(coreerrors.ErrEmailIsNotValid)
	}

	if len(cu.Password) < constants.MinimumPasswordLength {
		return errors.New(coreerrors.ErrMinimumPasswordLength)
	}

	if len(cu.Password) > constants.MaximumPasswordLength {
		return errors.New(coreerrors.ErrMaximumPasswordLength)
	}

	return nil
}
