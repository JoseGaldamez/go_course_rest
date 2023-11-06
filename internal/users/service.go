package users

import "log"

type (
	Service interface {
		Create(firstName, lastName, email, phone string) error
	}
	service struct{}
)

func NewService() Service {
	return &service{}
}

func (s service) Create(firstName, lastName, email, phone string) error {
	log.Println("Entrando al servicio de crear usuario")
	return nil
}
