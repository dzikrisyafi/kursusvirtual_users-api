package users

import (
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/utils/errors"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (login *LoginRequest) Validate() *errors.RestErr {
	if login.Username == "" && login.Password == "" {
		return errors.NewBadRequestError("invalid username and password")
	}
	if login.Username == "" {
		return errors.NewBadRequestError("invalid username")
	}
	if login.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
