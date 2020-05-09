package users

import (
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (login *LoginRequest) Validate() rest_errors.RestErr {
	if login.Username == "" && login.Password == "" {
		return rest_errors.NewBadRequestError("invalid username and password")
	}
	if login.Username == "" {
		return rest_errors.NewBadRequestError("invalid username")
	}
	if login.Password == "" {
		return rest_errors.NewBadRequestError("invalid password")
	}
	return nil
}
