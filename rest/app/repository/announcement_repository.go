package repository

import (
	"errors"
	"gin-api/rest/app/domain/dao"

	"gorm.io/gorm"
)

type AnnouncementRepository interface {
	Save(user *dao.User) (dao.User, error)
}

func NewAnnouncementRepository(db *gorm.DB) AnnouncementRepository {
	// db.AutoMigrate(&dao.User{})
	return &AnnouncementRepositoryImpl{
		db: db,
	}
}

type AnnouncementRepositoryImpl struct {
	db *gorm.DB
}

func (u *AnnouncementRepositoryImpl) Save(user *dao.User) (dao.User, error) {
	if user.Id == "" {
		return dao.User{}, errors.New("Invalid key")
	}
	if err := u.db.Save(user).Error; err != nil {
		return dao.User{}, err
	}
	err := u.db.Model(&dao.User{}).Preload("Role").Find(&user).Error
	return *user, err
}