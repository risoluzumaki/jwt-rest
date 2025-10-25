package user

import "errors"

type UserRepositoryInterface interface {
	GetByID(id int) (*User, error)
	GetByEmail(email string) (*User, error)
}

type UserRepository struct {
	user []User
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{
		user: []User{
			{
				ID:       1,
				Name:     "Test",
				Username: "testuser",
				Email:    "user1@test.com",
				Password: "passworduser1",
			},
		},
	}
}

func (r *UserRepository) GetByID(id int) (*User, error) {
	for _, user := range r.user {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) GetByEmail(email string) (*User, error) {
	for _, user := range r.user {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
