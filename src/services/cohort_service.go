package services

import (
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/cohort"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	CohortService cohortServiceInterface = &cohortService{}
)

type cohortService struct{}

type cohortServiceInterface interface {
	CreateCohort(cohort.Cohort) (*cohort.Cohort, rest_errors.RestErr)
	GetCohort(int) (*cohort.Cohort, rest_errors.RestErr)
	GetAllCohort() (cohort.Cohorts, rest_errors.RestErr)
	UpdateCohort(cohort.Cohort) (*cohort.Cohort, rest_errors.RestErr)
	DeleteCohort(int) rest_errors.RestErr
}

func (s *cohortService) CreateCohort(cohort cohort.Cohort) (*cohort.Cohort, rest_errors.RestErr) {
	if err := cohort.Validate(); err != nil {
		return nil, err
	}

	if err := cohort.Save(); err != nil {
		return nil, err
	}

	return &cohort, nil
}

func (s *cohortService) GetCohort(cohortID int) (*cohort.Cohort, rest_errors.RestErr) {
	dao := &cohort.Cohort{ID: cohortID}
	if err := dao.Get(); err != nil {
		return nil, err
	}

	return dao, nil
}

func (s *cohortService) GetAllCohort() (cohort.Cohorts, rest_errors.RestErr) {
	dao := &cohort.Cohort{}
	return dao.GetAll()
}

func (s *cohortService) UpdateCohort(cohort cohort.Cohort) (*cohort.Cohort, rest_errors.RestErr) {
	current, err := s.GetCohort(cohort.ID)
	if err != nil {
		return nil, err
	}

	if err := cohort.Validate(); err != nil {
		return nil, err
	}

	current.Name = cohort.Name
	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s cohortService) DeleteCohort(cohortID int) rest_errors.RestErr {
	dao := &cohort.Cohort{ID: cohortID}
	return dao.Delete()
}
