package model

import "gorm.io/gorm"

// 定义一个Model
type User struct {
	gorm.Model
	Name      string     `gorm:"type:varchar(20);not null"`
	Telephone string     `gorm:"varchar(11);not null;unique"`
	Password  string     `gorm:"size:255;not null"`
	Durations []Duration `gorm:"foreignKey:Tel;references:Telephone"`
}

type SimpleUser struct {
	Username  string
	TimeTotal int
}