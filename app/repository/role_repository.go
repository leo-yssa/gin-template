package repository

import (
	"gin-api/app/domain/dao"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Init() error
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	db.AutoMigrate(&dao.Role{}, &dao.User{})
	repo := &RoleRepositoryImpl{
		db: db,
	}
	repo.Init()
	return repo
}

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func (u *RoleRepositoryImpl) Init() error {
	u.db.Create(&dao.Role{
		ID: 1,
		Role: "sys",
	})
	u.db.Create(&dao.Role{
		ID: 2,
		Role: "oper",
	})
	u.db.Create(&dao.Role{
		ID: 3,
		Role: "user",
	})
	return nil
}