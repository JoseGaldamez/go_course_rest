package users

import (
	"log"
	"net/http"

	"gorm.io/gorm"
)

type (
	// ===================================== Endpoint =========================================
	Controller func(w http.ResponseWriter, r *http.Request)
	Endpoints  struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}
	CreateRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}
	UpdateRequest struct {
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Email     *string `json:"email"`
		Phone     *string `json:"phone"`
	}
	// ===================================== Error =========================================
	ErrorRequest struct {
		Error string `json:"error"`
	}
	// ===================================== Service =========================================
	Service interface {
		Create(firstName, lastName, email, phone string) (*User, error)
		GetAll() ([]User, error)
		Get(id string) (*User, error)
		Delete(id string) error
		Update(id string, firstName *string, lastName *string, email *string, phone *string) (map[string]interface{}, error)
	}
	UserService struct {
		repository Repository
		logger     *log.Logger
	} // Object which is going to include the interface

	// ===================================== Repository =========================================
	Repository interface {
		Create(user *User) error
		GetAll() ([]User, error)
		Get(id string) (*User, error)
		Delete(id string) error
		Update(id string, firstName *string, lastName *string, email *string, phone *string) (map[string]interface{}, error)
	}
	UserRepository struct {
		db     *gorm.DB
		logger *log.Logger
	}
)
