package database

import "github.com/harlancleiton/go-products/internal/entity"

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
