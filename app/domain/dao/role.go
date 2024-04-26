package dao

import "time"

type Role struct {
	ID   		int    		`gorm:"column:id; primary_key; not null" json:"id"`
	Role 		string 		`gorm:"column:role" json:"role"`
	CreatedAt   time.Time   `gorm:"column:createdAt; autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time   `gorm:"column:updatedAt; autoCreateTime" json:"updatedAt"`
	DeletedAt   time.Time   `gorm:"column:deletedAt; autoCreateTime" json:"deletedAt"`
}
