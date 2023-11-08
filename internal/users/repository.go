package users

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB, log *log.Logger) Repository {
	return &UserRepository{db: db, logger: log}
}

func (repo UserRepository) Create(user *User) error {
	repo.logger.Println("[!] Getting into Create User Method on repository layer")

	// create a id for the new user
	user.ID = uuid.New().String()

	// creating user on database
	result := repo.db.Create(user)
	if result.Error != nil {
		repo.logger.Println(result.Error)
		return result.Error
	}

	repo.logger.Println("[!] User was created with ID: " + user.ID)
	return nil
}

func (repository UserRepository) GetAll() ([]User, error) {
	repository.logger.Println("[!] Getting All Users on repository layer")

	var users []User

	result := repository.db.Model(&users).Order("created_at desc").Find(&users)
	if result.Error != nil {
		repository.logger.Println(result.Error)
		return nil, result.Error
	}

	return users, nil
}

func (repository UserRepository) Get(id string) (*User, error) {
	repository.logger.Println("[!] Getting One Users on repository layer")
	user := User{ID: id}

	result := repository.db.First(&user)
	if result.Error != nil {
		repository.logger.Println(result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func (repository UserRepository) Delete(id string) error {
	repository.logger.Println("[!] Delete User on repository layer")
	user := User{ID: id}

	result := repository.db.Delete(&user)
	if result.Error != nil {
		repository.logger.Println(result.Error)
		return result.Error
	}

	return nil
}

func (repository UserRepository) Update(id string, firstName *string, lastName *string, email *string, phone *string) (map[string]interface{}, error) {
	repository.logger.Println("[!] Update user on repository layer")
	values := make(map[string]interface{})
	if firstName != nil {
		values["first_name"] = firstName
	}
	if lastName != nil {
		values["last_name"] = lastName
	}
	if email != nil {
		values["email"] = email
	}
	if phone != nil {
		values["phone"] = phone
	}

	result := repository.db.Model(&User{}).Where("id = ?", id).Updates(values)
	if result.Error != nil {
		return nil, result.Error
	}

	return values, nil

}
