package submenu

import (
	"html"
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type SubMenu struct {
	ID       int    `json:"id"`
	MenuID   int    `json:"menu_id"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	Icon     string `json:"icon"`
	IsActive bool   `json:"is_active"`
}

type SubMenus []SubMenu

func (submenu SubMenu) Validate(isActive int) rest_errors.RestErr {
	submenu.Title = html.EscapeString(strings.TrimSpace(submenu.Title))
	submenu.URL = html.EscapeString(strings.TrimSpace(submenu.URL))
	submenu.Icon = html.EscapeString(strings.TrimSpace(submenu.Icon))

	if submenu.Title == "" {
		return rest_errors.NewBadRequestError("invalid submenu title")
	}

	if submenu.URL == "" {
		return rest_errors.NewBadRequestError("invalid submenu URL")
	}

	if submenu.Icon == "" {
		return rest_errors.NewBadRequestError("invalid submenu icon")
	}

	if isActive < 0 || isActive > 1 {
		return rest_errors.NewBadRequestError("invalid submenu status")
	}

	return nil
}
