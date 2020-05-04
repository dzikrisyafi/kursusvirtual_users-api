package services

import (
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/users"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/utils/crypto_utils"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/utils/date_utils"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	GetAllUser() (users.Users, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
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

func (s *usersService) GetAllUser() (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.GetAllUser()
}

func (s *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
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

func (s *usersService) DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(req users.LoginRequest) (*users.User, *errors.RestErr) {
	user := &users.User{Username: req.Username}

	// check if username and password is not empty
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

	// check username and password
	if err := dao.FindByUsernameAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
