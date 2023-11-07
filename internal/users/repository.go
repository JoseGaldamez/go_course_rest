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
