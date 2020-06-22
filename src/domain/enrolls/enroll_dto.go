package enrolls

import (
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type Course struct {
	CourseID int          `json:"course_id"`
	Users    []CourseUser `json:"users"`
}

type CourseUser struct {
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
}

type Enroll struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	CourseID int `json:"course_id"`
	CohortID int `json:"cohort_id"`
}

func (enroll Enroll) Validate() rest_errors.RestErr {
	if enroll.UserID <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}

	if enroll.CourseID <= 0 {
		return rest_errors.NewBadRequestError("invalid course id")
	}

	if enroll.CohortID <= 0 {
		return rest_errors.NewBadRequestError("invalid cohort id")
	}

	return nil
}
