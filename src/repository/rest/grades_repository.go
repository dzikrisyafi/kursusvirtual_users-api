package rest

import (
	"errors"
	"fmt"
	"time"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/golang-restclient/rest"
)

var (
	GradesRepository gradesRepositoryInterface = &gradesRepository{}
	gradesRestClient                           = rest.RequestBuilder{
		BaseURL: "http://localhost:8005",
		Timeout: 100 * time.Millisecond,
	}
)

type gradesRepository struct{}

type gradesRepositoryInterface interface {
	DeleteGrades(int, string) rest_errors.RestErr
}

func (r *gradesRepository) DeleteGrades(userID int, at string) rest_errors.RestErr {
	response := gradesRestClient.Delete(fmt.Sprintf("/internal/grades/users/%d?access_token=%s", userID, at))

	if response == nil || response.Response == nil {
		return rest_errors.NewInternalServerError("invalid rest client response when trying to delete grades", errors.New("rest client error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return rest_errors.NewInternalServerError("invalid error interface when trying to delete grades", err)
		}

		return apiErr
	}

	return nil
}
