package model

import (
	

	uuid "github.com/satori/go.uuid"
)

type User struct {
	Base
	Email string `json:"email"`
	Password string `json:"password"`
	EmailVerificationToken uuid.UUID
}

func CreateUser(user *User) (error, *User)  {
	result := db.Create(&user)

	if result.Error != nil {
		return result.Error, nil
	}

	return nil, user
}

func FetchUserByEmail(email string) ([]User, error) {
	var user []User
	result := db.Where("email = ?", email).Find(&user)
	
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}