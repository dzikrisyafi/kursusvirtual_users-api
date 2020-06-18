package users

import (
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"golang.org/x/net/html"
)

const (
	StatusActive = false
	DefaultImage = "default.jpg"
)

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Firstname    string `json:"firstname"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	RoleID       int    `json:"role_id"`
	DepartmentID int    `json:"department_id"`
	Salt         string `json:"salt"`
	Password     string `json:"password"`
	Image        string `json:"image"`
	Status       bool   `json:"status"`
	DateCreated  string `json:"date_created"`
}

type Users []User

func (user *User) Validate() rest_errors.RestErr {
	user.Username = html.EscapeString(strings.TrimSpace(strings.ToLower(user.Username)))
	user.Firstname = html.EscapeString(strings.TrimSpace(user.Firstname))
	user.Surname = html.EscapeString(strings.TrimSpace(user.Surname))
	user.Email = html.EscapeString(strings.TrimSpace(strings.ToLower(user.Email)))
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
