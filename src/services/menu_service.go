package services

import (
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/menu"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	MenuService menuServiceInterface = &menuService{}
)

type menuService struct{}

type menuServiceInterface interface {
	CreateMenu(menu.Menu) (*menu.Menu, rest_errors.RestErr)
	GetMenu(int) (*menu.Menu, rest_errors.RestErr)
	GetAllMenu() (menu.Menus, rest_errors.RestErr)
	GetAllMenuByRoleID(int) (menu.AccessMenus, rest_errors.RestErr)
	GetAllSubmenuByMenuID(*menu.AccessMenu) rest_errors.RestErr
	UpdateMenu(menu.Menu) (*menu.Menu, rest_errors.RestErr)
	DeleteMenu(int) rest_errors.RestErr
}

func (s *menuService) CreateMenu(menu menu.Menu) (*menu.Menu, rest_errors.RestErr) {
	if err := menu.Validate(); err != nil {
		return nil, err
	}

	if err := menu.Save(); err != nil {
		return nil, err
	}

	return &menu, nil
}

func (s *menuService) GetMenu(menuID int) (*menu.Menu, rest_errors.RestErr) {
	result := &menu.Menu{ID: menuID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *menuService) GetAllMenu() (menu.Menus, rest_errors.RestErr) {
	dao := &menu.Menu{}
	return dao.GetAll()
}

func (s *menuService) GetAllMenuByRoleID(roleID int) (menu.AccessMenus, rest_errors.RestErr) {
	dao := &menu.AccessMenu{}
	allMenu, err := dao.GetMenuByRoleID(roleID)
	if err != nil {
		return nil, err
	}

	result := make([]menu.AccessMenu, 0)
	for _, menu := range allMenu {
		if err := s.GetAllSubmenuByMenuID(&menu); err != nil {
			return nil, err
		}
		result = append(result, menu)
	}

	return result, nil
}

func (s *menuService) GetAllSubmenuByMenuID(allMenu *menu.AccessMenu) rest_errors.RestErr {
	dao := &menu.AccessSubMenu{}
	return dao.GetAllSubmenuByMenuID(allMenu)
}

func (s *menuService) UpdateMenu(menu menu.Menu) (*menu.Menu, rest_errors.RestErr) {
	current, err := s.GetMenu(menu.ID)
	if err != nil {
		return nil, err
	}

	if err := menu.Validate(); err != nil {
		return nil, err
	}

	current.Name = menu.Name
	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *menuService) DeleteMenu(menuID int) rest_errors.RestErr {
	dao := &menu.Menu{ID: menuID}
	return dao.Delete()
}
