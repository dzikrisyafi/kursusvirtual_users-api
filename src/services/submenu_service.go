package services

import (
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/submenu"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	SubMenuService submenuServiceInterface = &submenuService{}
)

type submenuService struct{}

type submenuServiceInterface interface {
	CreateSubMenu(submenu.SubMenu) (*submenu.SubMenu, rest_errors.RestErr)
	GetSubMenu(int) (*submenu.SubMenu, rest_errors.RestErr)
	GetAllSubMenu() (submenu.SubMenus, rest_errors.RestErr)
	UpdateSubMenu(bool, submenu.SubMenu) (*submenu.SubMenu, rest_errors.RestErr)
	DeleteSubMenu(int) rest_errors.RestErr
}

func (s *submenuService) CreateSubMenu(submenu submenu.SubMenu) (*submenu.SubMenu, rest_errors.RestErr) {
	var isActive int
	if submenu.IsActive {
		isActive = 1
	} else {
		isActive = 0
	}

	if err := submenu.Validate(isActive); err != nil {
		return nil, err
	}

	if err := submenu.Save(isActive); err != nil {
		return nil, err
	}

	return &submenu, nil
}

func (s *submenuService) GetSubMenu(submenuID int) (*submenu.SubMenu, rest_errors.RestErr) {
	result := &submenu.SubMenu{ID: submenuID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *submenuService) GetAllSubMenu() (submenu.SubMenus, rest_errors.RestErr) {
	result := &submenu.SubMenu{}
	return result.GetAll()
}

func (s *submenuService) UpdateSubMenu(isPartial bool, submenu submenu.SubMenu) (*submenu.SubMenu, rest_errors.RestErr) {
	current, err := s.GetSubMenu(submenu.ID)
	if err != nil {
		return nil, err
	}

	var isActive int
	if submenu.IsActive {
		isActive = 1
	} else {
		isActive = 0
	}

	if isPartial {
		if submenu.Title != "" {
			current.Title = submenu.Title
		}

		if submenu.URL != "" {
			current.URL = submenu.URL
		}

		if submenu.Icon != "" {
			current.Icon = submenu.Icon
		}

		if isActive == 0 || isActive == 1 {
			current.IsActive = submenu.IsActive
		}
	} else {
		if err := submenu.Validate(isActive); err != nil {
			return nil, err
		}

		current.Title = submenu.Title
		current.URL = submenu.URL
		current.Icon = submenu.Icon
	}

	if err := current.Update(isActive); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *submenuService) DeleteSubMenu(submenuID int) rest_errors.RestErr {
	result := &submenu.SubMenu{ID: submenuID}
	return result.Delete()
}
