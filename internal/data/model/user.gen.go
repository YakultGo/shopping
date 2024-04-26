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
	ID        int64      `gorm:"column:id;primaryKey" json:"id"`
	Name      string     `gorm:"column:name;not null" json:"name"`
	Address   *string    `gorm:"column:address" json:"address"`
	Telephone string     `gorm:"column:telephone;not null" json:"telephone"`
	Birthday  *time.Time `gorm:"column:birthday" json:"birthday"`
	DeleteAt  *time.Time `gorm:"column:delete_at" json:"delete_at"`
	Password  *string    `gorm:"column:password" json:"password"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}