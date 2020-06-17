package enrolls

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
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	CourseID int    `json:"course_id"`
	Cohort   string `json:"cohort"`
}
