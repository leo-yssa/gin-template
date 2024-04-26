package dao

import "time"

type User struct {
	Id 			string 		`gorm:"column:id; primaryKey" json:"id"`
	Email       string  	`gorm:"column:email; not null" binding:"required" json:"email"`           
	Name        string  	`gorm:"column:name; not null" binding:"required" json:"name"`
	Activated 	bool		`gorm:"column:activated; default:true" json:"activated"`
	RoleID   	int    		`gorm:"column:role_id;not null" json:"role_id"`
	Role     	Role   		`gorm:"foreignKey:RoleID;references:ID" json:"role"`
	CreatedAt   time.Time   `gorm:"column:createdAt; autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time   `gorm:"column:updatedAt; autoCreateTime" json:"updatedAt"`
}