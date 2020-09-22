package services

import (
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/departments"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	DepartmentsService departmentsServiceInterface = &departmentsService{}
)

type departmentsService struct{}

type departmentsServiceInterface interface {
	CreateDepartment(departments.Department) (*departments.Department, rest_errors.RestErr)
	GetDepartment(int) (*departments.Department, rest_errors.RestErr)
	GetAllDepartment() (departments.Departments, rest_errors.RestErr)
	UpdateDepartment(departments.Department) (*departments.Department, rest_errors.RestErr)
	DeleteDepartment(int) rest_errors.RestErr
}

func (s *departmentsService) CreateDepartment(department departments.Department) (*departments.Department, rest_errors.RestErr) {
	if err := department.Validate(); err != nil {
		return nil, err
	}

	if err := department.Save(); err != nil {
		return nil, err
	}

	return &department, nil
}

func (s *departmentsService) GetDepartment(departmentID int) (*departments.Department, rest_errors.RestErr) {
	result := &departments.Department{ID: departmentID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *departmentsService) GetAllDepartment() (departments.Departments, rest_errors.RestErr) {
	dao := &departments.Department{}
	return dao.GetAll()
}

func (s *departmentsService) UpdateDepartment(department departments.Department) (*departments.Department, rest_errors.RestErr) {
	current, err := s.GetDepartment(department.ID)
	if err != nil {
		return nil, err
	}

	if err := department.Validate(); err != nil {
		return nil, err
	}

	current.Name = department.Name
	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *departmentsService) DeleteDepartment(departmentID int) rest_errors.RestErr {
	dao := &departments.Department{ID: departmentID}
	return dao.Delete()
}
