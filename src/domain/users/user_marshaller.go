package users

import "encoding/json"

type PublicUser struct {
	ID          int64  `json:"id"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
}

type PrivateUser struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Firstname    string `json:"firstname"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	RoleID       int64  `json:"role_id"`
	DepartmentID int64  `json:"department_id"`
	Image        string `json:"image"`
	Status       string `json:"status"`
	DateCreated  string `json:"date_created"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			Status:      user.Status,
			DateCreated: user.DateCreated,
		}
	}

	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
