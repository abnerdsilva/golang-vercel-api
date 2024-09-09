package repository

import "errors"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type IUserRepository interface {
	FindAll() ([]User, error)
	FindByID(id int) (User, error)
}

// UserRepository é a estrutura para interagir com os dados dos usuários
type UserRepository struct {
}

// NewUserRepository cria um novo repositório com alguns dados iniciais
func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

var users = []User{
	{ID: 1, Name: "Alice", Age: 30},
	{ID: 2, Name: "Bob", Age: 25},
}

func (u UserRepository) FindAll() ([]User, error) {
	return users, nil
}

func (u UserRepository) FindByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}
