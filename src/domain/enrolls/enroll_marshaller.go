package enrolls

type PublicEnroll struct {
	CourseID int64 `json:"course_id"`
}

func (enroll Enroll) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicEnroll{
			CourseID: enroll.ID,
		}
	}
	return enroll
}
