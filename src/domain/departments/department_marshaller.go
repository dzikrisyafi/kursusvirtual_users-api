package departments

type PublicDepartment struct {
	ID int `json:"id"`
}

func (departments Departments) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(departments))
	for index, department := range departments {
		result[index] = department.Marshall(isPublic)
	}

	return result
}

func (department Department) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicDepartment{
			ID: department.ID,
		}
	}

	return department
}
