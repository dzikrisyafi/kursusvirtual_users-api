package roles

import (
	"html"
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type Role struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Roles []Role

func (role *Role) Validate() rest_errors.RestErr {
	role.Name = html.EscapeString(strings.TrimSpace(role.Name))
	if role.Name == "" {
		return rest_errors.NewBadRequestError("invalid role name")
	}

	return nil
}
