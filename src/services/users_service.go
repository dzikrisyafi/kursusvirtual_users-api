package services

import (
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/users"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/repository/rest"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/crypto_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/date_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, rest_errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, rest_errors.RestErr)
	GetAllUser() (users.Users, rest_errors.RestErr)
	GetUser(int) (*users.User, rest_errors.RestErr)
	DeleteUser(int, string) rest_errors.RestErr
	SearchUser(string) (users.Users, rest_errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, rest_errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Salt = crypto_utils.SaltText()
	user.Password = crypto_utils.GetPasswordHash(user.Password, user.Salt)
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) GetAllUser() (users.Users, rest_errors.RestErr) {
	dao := &users.User{}
	return dao.GetAllUser()
}

func (s *usersService) GetUser(userID int) (*users.User, rest_errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, rest_errors.RestErr) {
	current, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.Username != "" {
			current.Username = user.Username
		}

		if user.Firstname != "" {
			current.Firstname = user.Firstname
		}

		if user.Surname != "" {
			current.Surname = user.Surname
		}

		if user.Email != "" {
			current.Email = user.Email
		}

		if user.Password != "" {
			current.Salt = crypto_utils.SaltText()
			current.Password = crypto_utils.GetPasswordHash(user.Password, current.Salt)
		}

		if user.Image != "" {
			current.Image = user.Image
		}
	} else {
		if err := user.Validate(); err != nil {
			return nil, err
		}

		current.Username = user.Username
		current.Firstname = user.Firstname
		current.Surname = user.Surname
		current.Email = user.Email
		current.Salt = crypto_utils.SaltText()
		current.Password = crypto_utils.GetPasswordHash(user.Password, current.Salt)
		current.Image = user.Image
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userID int, at string) rest_errors.RestErr {
	user := &users.User{ID: userID}

	if err := rest.GradesRepository.DeleteGrades(userID, at); err != nil {
		return err
	}

	return user.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, rest_errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(req users.LoginRequest) (*users.User, rest_errors.RestErr) {
	user := &users.User{Username: req.Username}

	// validate username and password
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// get salt user by username
	if err := user.GetSalt(); err != nil {
		return nil, err
	}

	dao := &users.User{
		Username: req.Username,
		Password: crypto_utils.GetPasswordHash(req.Password, user.Salt),
	}

	// get user by username and password
	if err := dao.FindByUsernameAndPassword(); err != nil {
		return nil, err
	}

	return dao, nil
}
