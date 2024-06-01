package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"type:varchar(255);uniqueIndex"`
	PasswordHashed string `gorm:"type:varchar(255);not null"`
}

func (User) TableName() string {
	return "user"
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetByEmail(db *gorm.DB, email string) (*User, error) {
	user := &User{}
	err := db.Where("email=%s", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
