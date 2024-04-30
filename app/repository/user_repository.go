package repository

import (
	"gin-api/app/domain/dao"

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
	if err := u.db.Save(user).Error; err != nil {
		return dao.User{}, err
	}
	return *user, nil
}