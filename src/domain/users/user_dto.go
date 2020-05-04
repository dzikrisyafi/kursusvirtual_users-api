package users

import (
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_users-api/src/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Firstname    string `json:"firstname"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	RoleID       int64  `json:"role_id"`
	DepartmentID int64  `json:"department_id"`
	Salt         string `json:"salt"`
	Password     string `json:"password"`
	Image        string `json:"image"`
	Status       string `json:"status"`
	DateCreated  string `json:"date_created"`
}

type UserRole struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type UserDepartment struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.Username = strings.TrimSpace(strings.ToLower(user.Username))
	user.Firstname = strings.TrimSpace(user.Firstname)
	user.Surname = strings.TrimSpace(user.Surname)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	if user.Username == "" {
		return errors.NewBadRequestError("invalid username")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}
