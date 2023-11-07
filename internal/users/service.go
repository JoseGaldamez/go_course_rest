package users

import (
	"log"

	"gorm.io/gorm"
)

func NewService(db *gorm.DB, log *log.Logger) Service {

	repository := NewRepository(db, log)

	return &UserService{
		repo:   repository,
		logger: log,
	}
}

func (service UserService) Create(firstName, lastName, email, phone string) (*User, error) {
	service.logger.Println("[!] Getting into Create User Method on service layer")
	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	err := service.repo.Create(&user)
	if err != nil {
		service.logger.Println("[x] Something went wrong with the respository response")
		return nil, err
	}
	return &user, nil
}
