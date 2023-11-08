package users

import (
	"log"

	"gorm.io/gorm"
)

func NewService(db *gorm.DB, log *log.Logger) Service {

	repository := NewRepository(db, log)

	return &UserService{
		repository: repository,
		logger:     log,
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

	err := service.repository.Create(&user)
	if err != nil {
		service.logger.Println("[x] Something went wrong with the respository response")
		return nil, err
	}
	return &user, nil
}

func (service UserService) GetAll() ([]User, error) {
	service.logger.Println("[!] Getting All User on service layer")
	users, err := service.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service UserService) Get(id string) (*User, error) {
	service.logger.Println("[!] Getting One User on service layer")
	user, err := service.repository.Get(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service UserService) Delete(id string) error {
	service.logger.Println("[!] Deleting User on service layer")
	err := service.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (service UserService) Update(id string, firstName *string, lastName *string, email *string, phone *string) (map[string]interface{}, error) {
	service.logger.Println("[!] Update User on service layer")
	values, err := service.repository.Update(id, firstName, lastName, email, phone)
	if err != nil {
		return nil, err
	}

	return values, nil
}
