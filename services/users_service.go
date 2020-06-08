package services

import (
	"github.com/posol/bookstore_users_api/domain/users"
	"github.com/posol/bookstore_users_api/utils/crypto"
	"github.com/posol/bookstore_users_api/utils/dates"
	"github.com/posol/bookstore_users_api/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestError)
	CreateUser(users.User) (*users.User, *errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *errors.RestError)
	DeleteUser(int64) *errors.RestError
	SearchUser(string) (users.Users, *errors.RestError)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestError)
}

type userService struct {
}

func (u *userService) GetUser(userId int64) (*users.User, *errors.RestError) {
	user := users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = dates.GetNowDBFormat()
	user.Password = crypto.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	current, err := u.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (u *userService) DeleteUser(userId int64) *errors.RestError {
	user := &users.User{Id: userId}
	return user.Delete()

}

func (u *userService) SearchUser(status string) (users.Users, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (u *userService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestError) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
