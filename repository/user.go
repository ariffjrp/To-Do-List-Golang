package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var data entity.User
	dataTerakhir := r.db.WithContext(ctx).Where("id = ?", id).Model(&entity.User{}).Find(&data)
	Alert := dataTerakhir.Error
	if Alert != nil {
		if errors.Is(Alert, gorm.ErrRecordNotFound) {
		return entity.User{}, nil
	  }else{
		return entity.User{}, Alert
	  }
   }

	return data, nil // TODO: replace this
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var data entity.User
	dataTerakhir := r.db.WithContext(ctx).Where("email = ?", email).Model(&entity.User{}).Find(&data)
	Alert := dataTerakhir.Error
	if Alert != nil {
		if errors.Is(Alert, gorm.ErrRecordNotFound) {
		return entity.User{}, nil
	  }else{
		return entity.User{}, Alert
	  }
   }

	return data, nil // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	dataTerakhir := r.db.WithContext(ctx).Create(&user)
	Alert := dataTerakhir.Error
	if Alert != nil {
		return entity.User{}, Alert
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	dataTerakhir := r.db.WithContext(ctx).Model(&user).
	Updates(user)
	Alert := dataTerakhir.Error
	if Alert != nil {
		return entity.User{}, Alert
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	DataTerakhir := r.db.WithContext(ctx).Where("id = ?", id).
		Delete(&entity.User{})
	Alert := DataTerakhir.Error
	if Alert != nil {
		return Alert
	}
	return nil // TODO: replace this
}
