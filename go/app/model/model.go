package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Usertoken string `gorm:"not null"`
}

// ユーザーチャットのモデル
type Userchatroom struct {
	gorm.Model
	Firstuserid  uint
	Seconduserid uint
}

type UserMessage struct {
	gorm.Model
	UserchatroomID uint
	UserID         uint
	Message        string
}
