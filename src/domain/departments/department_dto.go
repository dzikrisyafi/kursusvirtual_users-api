package departments

import (
	"html"
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Departments []Department

func (department Department) Validate() rest_errors.RestErr {
	department.Name = html.EscapeString(strings.TrimSpace(department.Name))
	if department.Name == "" {
		return rest_errors.NewBadRequestError("invalid department name")
	}

	return nil
}
