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
	GetUsersByCourseID(int64) (*enrolls.Course, rest_errors.RestErr)
	CreateEnroll(enrolls.Enroll) (*enrolls.Enroll, rest_errors.RestErr)
}

func (s *enrollsService) GetUsersByCourseID(courseID int64) (*enrolls.Course, rest_errors.RestErr) {
	result := &enrolls.Course{CourseID: courseID}
	if err := result.FindUserByCourseID(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *enrollsService) CreateEnroll(req enrolls.Enroll) (*enrolls.Enroll, rest_errors.RestErr) {
	dao := &enrolls.Enroll{
		UserID:   req.UserID,
		CourseID: req.CourseID,
		Cohort:   req.Cohort,
	}

	if err := dao.Save(); err != nil {
		return nil, err
	}

	return dao, nil
}
