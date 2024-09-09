package repository

import (
	"errors"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// UserRepository é a estrutura para interagir com os dados dos usuários
type UserRepository struct {
	users []User
}

// NewUserRepository cria um novo repositório com alguns dados iniciais
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []User{
			{ID: 1, Name: "Alice", Age: 30},
			{ID: 2, Name: "Bob", Age: 25},
		},
	}
}

// FindAll retorna todos os usuários
func (repo *UserRepository) FindAll() []User {
	return repo.users
}

// FindByID procura um usuário pelo ID
func (repo *UserRepository) FindByID(id int) (User, error) {
	for _, user := range repo.users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}
