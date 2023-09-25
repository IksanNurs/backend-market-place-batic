// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"e-commerce/helpers"
	"time"

	"github.com/jinzhu/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID           int32  `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`
	Name         string `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Email        string `gorm:"column:email;type:varchar(50);not null" json:"email"`
	Phone        string `gorm:"column:phone;type:varchar(50);not null" json:"phone"`
	PasswordHash string `gorm:"column:password_hash;type:varchar(255);not null" json:"-"`
	AuthKey      string `gorm:"column:auth_key;type:text;not null" json:"-"`
	IsSales      int32  `gorm:"column:is_sales;type:int(11);not null" json:"is_sales"`
	CreatedAt    int32  `gorm:"column:created_at;type:int(11);not null" json:"created_at"`
	UpdatedAt    int32  `gorm:"column:updated_at;type:int(11);not null" json:"updated_at"`
}

type InputUser struct {
	Name         string `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Email        string `gorm:"column:email;type:varchar(50);not null" json:"email" binding:"required"`
	Phone        string `gorm:"column:phone;type:varchar(50);not null" json:"phone"`
	PasswordHash string `gorm:"column:password_hash;type:varchar(255);not null" json:"password_hash" binding:"required"`
}

type UpdateUser struct {
	Name         string `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Email        string `gorm:"column:email;type:varchar(50);not null" json:"email" binding:"required"`
	Phone        string `gorm:"column:phone;type:varchar(50);not null" json:"phone"`
}

type InputUser1 struct {
	Phone        string `gorm:"column:phone;type:varchar(50);not null" json:"phone"`
	Email        string `gorm:"column:email;type:varchar(50);not null" json:"email" binding:"required"`
	PasswordHash string `gorm:"column:password_hash;type:varchar(255);not null" json:"password_hash" binding:"required"`
	CreatedAt    int32  `gorm:"column:created_at;type:int(11);not null" json:"created_at"`
}




func (i *User) BeforeCreate(scope *gorm.Scope) error {
    now := int32(time.Now().Unix())
    i.CreatedAt = now
	 i.PasswordHash= helpers.HassPass(i.PasswordHash)

    return nil
}

func (i *InputUser1) BeforeCreate(scope *gorm.Scope) error {
    now := int32(time.Now().Unix())
    i.CreatedAt = now

    return nil
}



// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

func (*InputUser1) TableName() string {
	return TableNameUser
}


func (*UpdateUser) TableName() string {
	return TableNameUser
}


