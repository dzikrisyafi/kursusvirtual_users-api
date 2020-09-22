package menu

import (
	"html"
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type Menu struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Menus []Menu

func (menu Menu) Validate() rest_errors.RestErr {
	menu.Name = html.EscapeString(strings.TrimSpace(menu.Name))
	if menu.Name == "" {
		return rest_errors.NewBadRequestError("invalid menu name")
	}

	return nil
}
