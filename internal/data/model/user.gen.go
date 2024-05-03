// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Address   string    `gorm:"column:address" json:"address"`
	Telephone string    `gorm:"column:telephone;not null" json:"telephone"`
	Birthday  time.Time `gorm:"column:birthday;default:null" json:"birthday"`
	Password  string    `gorm:"column:password;not null" json:"password"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
