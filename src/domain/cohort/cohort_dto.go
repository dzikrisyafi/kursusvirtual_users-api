package cohort

import (
	"html"
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type Cohort struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Cohorts []Cohort

func (cohort Cohort) Validate() rest_errors.RestErr {
	cohort.Name = html.EscapeString(strings.TrimSpace(cohort.Name))
	if cohort.Name == "" {
		return rest_errors.NewBadRequestError("invalid cohort name")
	}

	return nil
}
