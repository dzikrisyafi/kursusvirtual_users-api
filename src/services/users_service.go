package services

import (
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/users"
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
	LoginUser(users.LoginRequest) (*users.User, rest_errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Salt = crypto_utils.SaltText()
	user.Password = crypto_utils.GetPasswordHash(user.Password, user.Salt)
	user.Image = users.DefaultImage
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	if err := user.Save(0); err != nil {
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

	var status int
	if user.Status {
		status = 1
	} else {
		status = 0
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

		if user.RoleID > 0 {
			current.RoleID = user.RoleID
		}

		if user.DepartmentID > 0 {
			current.DepartmentID = user.DepartmentID
		}

		if user.Password != "" {
			current.Salt = crypto_utils.SaltText()
			current.Password = crypto_utils.GetPasswordHash(user.Password, current.Salt)
		}

		if user.Image != "" {
			current.Image = user.Image
		}

		if status == 0 || status == 1 {
			current.Status = user.Status
		}
	} else {
		if err := user.Validate(); err != nil {
			return nil, err
		}

		current.Username = user.Username
		current.Firstname = user.Firstname
		current.Surname = user.Surname
		current.Email = user.Email
		current.RoleID = user.RoleID
		current.DepartmentID = user.DepartmentID
		current.Salt = crypto_utils.SaltText()
		current.Password = crypto_utils.GetPasswordHash(user.Password, current.Salt)
		current.Image = users.DefaultImage
	}

	if err := current.Update(status); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *usersService) DeleteUser(userID int, at string) rest_errors.RestErr {
	user := &users.User{ID: userID}

	// if err := rest.GradesRepository.DeleteGrades(userID, at); err != nil {
	// 	return err
	// }

	return user.Delete()
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
