package enrolls

import (
	"errors"
	"fmt"

	"github.com/dzikrisyafi/kursusvirtual_users-api/src/datasources/mysql/users_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

const (
	queryInsertEnroll      = `INSERT INTO enrolls (user_id, course_id, cohort_id) VALUES (?, ?, ?);`
	queryGetUserByCourseID = `SELECT user_id, username, firstname, surname, email FROM enrolls INNER JOIN users ON user_id=users.id WHERE course_id=?;`
	queryDeleteEnroll      = `DELETE FROM enrolls WHERE id=?;`
)

func (enroll *Enroll) Save() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryInsertEnroll)
	if err != nil {
		logger.Error("error when trying to prepare save enroll statement", err)
		return rest_errors.NewInternalServerError("error when trying to save enroll", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(enroll.UserID, enroll.CourseID, enroll.Cohort)
	if saveErr != nil {
		logger.Error("error when trying to save enroll", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save enroll", errors.New("database error"))
	}

	enrollID, err := insertResult.LastInsertId()
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to get last insert id after creating a new role", errors.New("database error"))
	}
	enroll.ID = int(enrollID)

	return nil
}

func (course *Course) GetUserByCourseID() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryGetUserByCourseID)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to get users by course id", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(course.CourseID)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to get users by course id", errors.New("database error"))
	}
	defer rows.Close()

	var user CourseUser
	for rows.Next() {
		if err := rows.Scan(&user.UserID, &user.Username, &user.Firstname, &user.Surname, &user.Email); err != nil {
			return rest_errors.NewInternalServerError("error when trying to get users by course id", errors.New("database error"))
		}

		course.Users = append(course.Users, user)
	}

	if len(course.Users) == 0 {
		return rest_errors.NewNotFoundError(fmt.Sprintf("no users matching given course id %d", course.CourseID))
	}

	return nil
}

func (enroll *Enroll) DeleteEnroll() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryDeleteEnroll)
	if err != nil {
		logger.Error("error when trying to prepare delete user enroll by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete user enroll", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err := stmt.Exec(enroll.ID); err != nil {
		logger.Error("error when trying to delete user enroll by id", err)
		return rest_errors.NewInternalServerError("error when trying to delete user enroll", errors.New("database error"))
	}

	return nil
}
