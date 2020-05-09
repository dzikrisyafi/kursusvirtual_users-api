package users

import (
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
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

type Users []User

func (user *User) Validate() rest_errors.RestErr {
	user.Username = strings.TrimSpace(strings.ToLower(user.Username))
	user.Firstname = strings.TrimSpace(user.Firstname)
	user.Surname = strings.TrimSpace(user.Surname)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	if user.Username == "" {
		return rest_errors.NewBadRequestError("invalid username")
	}
	if user.Password == "" {
		if len(user.Password) < 8 {
			return rest_errors.NewBadRequestError("the password must have at least 8 characters")
		}
		return rest_errors.NewBadRequestError("invalid password")
	}
	return nil
}
