package usecase

import (
	"context"
	dto "go-template/dto/general"
	dtousecase "go-template/dto/general/usecase"
	"go-template/entity"
	"go-template/repository"
	appconstant "go-template/share/general/constant"
	errorapp "go-template/share/general/error"
	"go-template/share/general/util"
	"time"

	"github.com/go-errors/errors"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, cu dtousecase.CreateUserRequest) (dtousecase.CreateUserResponse, error)
	LoginUser(ctx context.Context, cu dtousecase.LoginUserRequest) (dtousecase.LoginUserResponse, error)
	ForgetPassword(ctx context.Context, r dtousecase.ForgetPasswordRequest) (dtousecase.ForgetPasswordResponse, error)
	ResetPassword(ctx context.Context, r dtousecase.ResetPasswordRequest) error
}

type userUsecase struct {
	userRepository repository.UserRepository

	authTokenGenerator   util.AuthTokenGenerator
	passwordHasher       util.PasswordHasher
	randomTokenGenerator util.RandomTokenGenerator
}

type UserUsecaseConfig struct {
	UserRepository       repository.UserRepository
	PasswordHasher       util.PasswordHasher
	AuthTokenGenerator   util.AuthTokenGenerator
	RandomTokenGenerator util.RandomTokenGenerator
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

func (uu *userUsecase) CreateUser(ctx context.Context, cu dtousecase.CreateUserRequest) (dtousecase.CreateUserResponse, error) {
	cures := dtousecase.CreateUserResponse{}

	if err := uu.checkCreateUserData(ctx, cu); err != nil {
		return cures, errorapp.ErrorHandling(err)
	}

	hashedPassword, err := uu.passwordHasher.GenerateHashFromPassword(cu.Password)

	if err != nil {
		return cures, errorapp.ErrorHandling(err)
	}

	u := entity.User{
		Name:     cu.Name,
		Email:    cu.Email,
		Password: hashedPassword,
	}

	u2, err := uu.userRepository.Create(ctx, u)

	if err != nil {
		return cures, errorapp.ErrorHandling(err)
	}

	cures.Name = u2.Name
	cures.Email = u2.Email

	return cures, nil
}

func (uu *userUsecase) checkCreateUserData(ctx context.Context, cu dtousecase.CreateUserRequest) error {
	if !util.NewDataValidator().IsEmailValid(cu.Email) {
		return errors.New(errorapp.ErrEmailIsNotValid)
	}

	if len(cu.Password) < appconstant.MinimumPasswordLength {
		return errors.New(errorapp.ErrMinimumPasswordLength)
	}

	if len(cu.Password) > appconstant.MaximumPasswordLength {
		return errors.New(errorapp.ErrMaximumPasswordLength)
	}

	return nil
}

func (uu *userUsecase) LoginUser(ctx context.Context, cu dtousecase.LoginUserRequest) (dtousecase.LoginUserResponse, error) {
	lures := dtousecase.LoginUserResponse{}

	if err := uu.checkLoginUserData(ctx, cu); err != nil {
		return lures, errorapp.ErrorHandling(err)
	}

	u, err := uu.userRepository.FindByEmail(ctx, cu.Email)

	if err != nil {
		return lures, errorapp.ErrorHandling(err)
	}

	if u.ID == 0 {
		return lures, errors.New(errorapp.ErrEmailNotFound)
	}

	match, err := uu.passwordHasher.CompareHashAndPassword(cu.Password, u.Password)

	if !match {
		return lures, errors.New(errorapp.ErrWrongPassword)
	}

	if err != nil {
		return lures, errorapp.ErrorHandling(err)
	}

	ad := dto.AuthData{
		ID: u.ID,
	}

	token, err := uu.authTokenGenerator.Encode(ad)

	if err != nil {
		return lures, errorapp.ErrorHandling(err)
	}

	lures.Token = token

	return lures, nil
}

func (uu *userUsecase) checkLoginUserData(ctx context.Context, cu dtousecase.LoginUserRequest) error {
	if !util.NewDataValidator().IsEmailValid(cu.Email) {
		return errors.New(errorapp.ErrEmailIsNotValid)
	}

	if len(cu.Password) < appconstant.MinimumPasswordLength {
		return errors.New(errorapp.ErrMinimumPasswordLength)
	}

	if len(cu.Password) > appconstant.MaximumPasswordLength {
		return errors.New(errorapp.ErrMaximumPasswordLength)
	}

	return nil
}

func (uu *userUsecase) ForgetPassword(ctx context.Context, r dtousecase.ForgetPasswordRequest) (dtousecase.ForgetPasswordResponse, error) {
	res := dtousecase.ForgetPasswordResponse{}

	if err := uu.checkForgetPasswordData(ctx, r); err != nil {
		return res, errorapp.ErrorHandling(err)
	}

	u, err := uu.userRepository.FindByEmail(ctx, r.Email)

	if err != nil {
		return res, errorapp.ErrorHandling(err)
	}

	if u.ID == 0 {
		return res, errors.New(errorapp.ErrEmailNotFound)
	}

	rt, err := uu.randomTokenGenerator.Generate(appconstant.ForgetPasswordTokenLength)

	if err != nil {
		return res, errorapp.ErrorHandling(err)
	}

	err = uu.userRepository.DeletePreviousResetPassword(ctx, int(u.ID))

	if err != nil {
		return res, errorapp.ErrorHandling(err)
	}

	urp := entity.UserResetPassword{
		UserId:    int(u.ID),
		Token:     rt,
		ExpiredAt: time.Now().Add(appconstant.ForgetPasswordExpiredDuration),
	}

	urp2, err := uu.userRepository.CreateUserForgetPassword(ctx, urp)

	if err != nil {
		return res, errorapp.ErrorHandling(err)
	}

	res.Token = urp2.Token
	res.ExpiredAt = urp2.ExpiredAt

	return res, nil
}

func (uu *userUsecase) checkForgetPasswordData(ctx context.Context, r dtousecase.ForgetPasswordRequest) error {
	if !util.NewDataValidator().IsEmailValid(r.Email) {
		return errors.New(errorapp.ErrEmailIsNotValid)
	}

	return nil
}

func (uu *userUsecase) ResetPassword(ctx context.Context, r dtousecase.ResetPasswordRequest) error {
	if err := uu.checkResetPasswordData(ctx, r); err != nil {
		return errorapp.ErrorHandling(err)
	}

	rt, err := uu.userRepository.GetResetPasswordTokenByToken(ctx, r.Token)

	if err != nil {
		return errorapp.ErrorHandling(err)
	}

	if rt.ID == 0 {
		return errors.New("token not found")
	}

	u, err := uu.userRepository.FindByEmail(ctx, r.Email)

	if err != nil {
		return errorapp.ErrorHandling(err)
	}

	if u.ID == 0 {
		return errors.New("user not found")
	}

	if rt.UserId != int(u.ID) {
		return errors.New("reset code is not registered for this account")
	}

	hp, err := uu.passwordHasher.GenerateHashFromPassword(r.NewPassword)

	if err != nil {
		return errorapp.ErrorHandling(err)
	}

	u.Password = hp

	_, err = uu.userRepository.UpdateUser(ctx, u)

	if err != nil {
		return errorapp.ErrorHandling(err)
	}

	_, err = uu.userRepository.DeleteUsedResetPassword(ctx, rt)

	if err != nil {
		return errorapp.ErrorHandling(err)
	}

	return nil
}

func (uu *userUsecase) checkResetPasswordData(ctx context.Context, r dtousecase.ResetPasswordRequest) error {
	if !util.NewDataValidator().IsEmailValid(r.Email) {
		return errors.New(errorapp.ErrEmailIsNotValid)
	}

	if len(r.NewPassword) < appconstant.MinimumPasswordLength {
		return errors.New(errorapp.ErrMinimumPasswordLength)
	}

	if len(r.Token) != appconstant.ForgetPasswordTokenLength {
		return errors.New(errorapp.ErrForgetPasswordTokenLength)
	}

	return nil
}
