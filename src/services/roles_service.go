package services

import (
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/roles"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	RolesService rolesServiceInterface = &rolesService{}
)

type rolesService struct{}

type rolesServiceInterface interface {
	GetRole(int64) (*roles.Role, rest_errors.RestErr)
	GetAllRole() (roles.Roles, rest_errors.RestErr)
	CreateRole(roles.Role) (*roles.Role, rest_errors.RestErr)
	UpdateRole(roles.Role) (*roles.Role, rest_errors.RestErr)
	DeleteRole(int64) rest_errors.RestErr
}

func (s *rolesService) GetRole(roleID int64) (*roles.Role, rest_errors.RestErr) {
	result := &roles.Role{ID: roleID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *rolesService) GetAllRole() (roles.Roles, rest_errors.RestErr) {
	dao := &roles.Role{}
	return dao.GetAllRole()
}

func (s *rolesService) CreateRole(role roles.Role) (*roles.Role, rest_errors.RestErr) {
	if err := role.Validate(); err != nil {
		return nil, err
	}

	if err := role.Save(); err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *rolesService) UpdateRole(role roles.Role) (*roles.Role, rest_errors.RestErr) {
	current, err := s.GetRole(role.ID)
	if err != nil {
		return nil, err
	}

	if err := role.Validate(); err != nil {
		return nil, err
	}

	current.Name = role.Name

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *rolesService) DeleteRole(roleID int64) rest_errors.RestErr {
	dao := &roles.Role{ID: roleID}
	return dao.Delete()
}
