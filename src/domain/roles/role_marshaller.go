package roles

type PublicRole struct {
	ID int `json:"id"`
}

func (roles Roles) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(roles))
	for index, role := range roles {
		result[index] = role.Marshall(isPublic)
	}

	return result
}

func (role Role) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicRole{
			ID: role.ID,
		}
	}

	return role
}
