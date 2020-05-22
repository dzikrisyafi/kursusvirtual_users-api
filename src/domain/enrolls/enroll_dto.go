package enrolls

type Course struct {
	CourseID int64        `json:"course_id"`
	Users    []CourseUser `json:"users"`
}

type CourseUser struct {
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
}

type Enroll struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	CourseID int64  `json:"course_id"`
	Cohort   string `json:"cohort"`
}
