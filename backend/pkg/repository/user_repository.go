package repository

import (
	"context"
	"errors"

	"github.com/SatriaAPN/my-e-wallet/backend/pkg/core"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, u core.User) (core.User, error)
	FindByEmail(ctx context.Context, email string) (core.User, error)
	FindById(ctx context.Context, userId int) (core.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type UserRepositoryConfig struct {
	Db *gorm.DB
}

func NewUserRepository(config UserRepositoryConfig) UserRepository {
	br := userRepository{}

	if config.Db != nil {
		br.db = config.Db
	}

	return &br
}

func (ur *userRepository) Create(ctx context.Context, u core.User) (core.User, error) {

	err := ur.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&u).Error

		if err != nil {
			return errors.New("error")
		}

		return nil
	})

	return u, err
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (core.User, error) {
	u := core.User{}

	err := ur.db.WithContext(ctx).Where("email = ?", email).First(&u).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u, errors.New("error")
	}

	return u, err
}

func (ur *userRepository) FindById(ctx context.Context, userId int) (core.User, error) {
	u := core.User{}

	err := ur.db.WithContext(ctx).Where("id = ?", userId).First(&u).Error

	return u, err
}
