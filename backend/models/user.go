package models

import (
	"errors"
	"time"

	"github.com/leo-yssa/gin-template/lib"
)

type User struct {
	Id string `gorm:"primaryKey"`
	Email        string        `gorm:"not null" binding:"required"`           
	Kind 		string		`gorm:"not null" binding:"required"`
	Name         string         `gorm:"not null" binding:"required"`
	// Email        *string        // A pointer to a string, allowing for null values
	// Age          uint8          // An unsigned 8-bit integer
	// Birthday     *time.Time     // A pointer to time.Time, can be null
	// MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	// ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	Activated bool				`gorm:"default:true"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoCreateTime"`
}

func (m *Model) NewUser() *User{
	return new(User)
}

func (m *Model) RegisterUser(_user *User) (*User, error){
	var user *User
	id, err := lib.Sha3Hash(_user.Email)
	if err != nil {
		return nil, err
	}
	_user.Id = id
	tx := m.DB.Find(&user, &User{
		Id: id,
	})
	if tx.Error != nil {
		return nil, tx.Error
	}
	if user.Email == _user.Email {
		return nil, errors.New("already registered")
	}

	tx = m.DB.Create(&_user)
	if tx.Error != nil {
		tx.Rollback()
		return nil, tx.Error
	}
	return _user, nil
}