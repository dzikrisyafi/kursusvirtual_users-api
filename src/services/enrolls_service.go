package services

import (
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	EnrollsService enrollsServiceInterface = &enrollsService{}
)

type enrollsService struct{}

type enrollsServiceInterface interface {
	GetUsersByCourseID(int) (*enrolls.Course, rest_errors.RestErr)
	CreateEnroll(enrolls.Enroll) (*enrolls.Enroll, rest_errors.RestErr)
	UpdateEnrollByUserIDAndCourseID(enrolls.Enroll) (*enrolls.Enroll, rest_errors.RestErr)
	DeleteEnroll(int) rest_errors.RestErr
}

func (s *enrollsService) GetUsersByCourseID(courseID int) (*enrolls.Course, rest_errors.RestErr) {
	result := &enrolls.Course{CourseID: courseID}
	if err := result.GetUserByCourseID(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *enrollsService) CreateEnroll(req enrolls.Enroll) (*enrolls.Enroll, rest_errors.RestErr) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	dao := &enrolls.Enroll{
		UserID:   req.UserID,
		CourseID: req.CourseID,
		CohortID: req.CohortID,
	}

	if err := dao.Save(); err != nil {
		return nil, err
	}

	return dao, nil
}

func (s *enrollsService) UpdateEnrollByUserIDAndCourseID(req enrolls.Enroll) (*enrolls.Enroll, rest_errors.RestErr) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	dao := &enrolls.Enroll{
		UserID:   req.UserID,
		CourseID: req.CourseID,
		CohortID: req.CohortID,
	}

	if err := dao.Update(); err != nil {
		return nil, err
	}

	return dao, nil
}

func (s *enrollsService) DeleteEnroll(enrollID int) rest_errors.RestErr {
	dao := &enrolls.Enroll{ID: enrollID}
	return dao.DeleteEnroll()
}
