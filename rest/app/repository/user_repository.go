package repository

import (
	"errors"
	"gin-api/rest/app/domain/dao"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *dao.User) (dao.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	// db.AutoMigrate(&dao.User{})
	return &UserRepositoryImpl{
		db: db,
	}
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u *UserRepositoryImpl) Save(user *dao.User) (dao.User, error) {
	if user.Id == "" {
		return dao.User{}, errors.New("Invalid key")
	}
	if err := u.db.Save(user).Error; err != nil {
		return dao.User{}, err
	}
	err := u.db.Model(&dao.User{}).Preload("Role").Find(&user).Error
	return *user, err
}