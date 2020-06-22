package cohort

type PublicCohort struct {
	ID int `json:"id"`
}

func (cohorts Cohorts) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(cohorts))
	for index, cohort := range cohorts {
		result[index] = cohort.Marshall(isPublic)
	}

	return result
}

func (cohort Cohort) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicCohort{
			ID: cohort.ID,
		}
	}

	return cohort
}
