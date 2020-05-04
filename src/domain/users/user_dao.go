package users

import (
	"fmt"
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_users-api/src/datasources/mysql/users_db"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/logger"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/utils/errors"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/utils/mysql_utils"
)

// query for database
const (
	queryGetUser               = `SELECT username, firstname, surname, email, role_id, department_id, salt, password, image, status, date_created FROM users WHERE id=?;`
	queryGetAllUser            = `SELECT id, username, firstname, surname, email, role_id, department_id, image, status, date_created FROM users;`
	queryInsertUser            = `INSERT INTO users(username, firstname, surname, email, role_id, department_id, salt, password, image, status, date_created) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	queryUpdateUser            = `UPDATE users SET username=?, firstname=?, surname=?, email=?, role_id=?, department_id=?, salt=?, password=?, image=? WHERE id=?;`
	queryDeleteUser            = `DELETE FROM users WHERE id=?;`
	queryFindUserByStatus      = `SELECT id, username, firstname, surname, email, role_id, department_id, image, status, date_created FROM users WHERE status=?;`
	queryFindByUserAndPassword = `SELECT id, username, firstname, surname, email, image, status, date_created FROM users WHERE username=? AND password=?;`
	queryGetSalt               = `SELECT salt FROM users WHERE username=?;`
)

// data access object method
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.Username, user.Firstname, user.Surname, user.Email, user.RoleID, user.DepartmentID, user.Salt, user.Password, user.Image, user.Status, user.DateCreated)
	if saveErr != nil {
		logger.Error("error when trying to save user", err)
		return mysql_utils.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.ID = userID
	return nil
}

func (user *User) GetAllUser() ([]User, *errors.RestErr) {
	stmt, err := users_db.DbConn().Prepare(queryGetAllUser)
	if err != nil {
		logger.Error("error when trying to prepare get all user statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, getErr := stmt.Query()
	if getErr != nil {
		logger.Error("error when trying to get all user", getErr)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Firstname, &user.Surname, &user.Email, &user.RoleID, &user.DepartmentID, &user.Image, &user.Status, &user.DateCreated); err != nil {
			logger.Error("error when trying to scan user rows into user struct", err)
			return nil, mysql_utils.ParseError(err)
		}

		result = append(result, *user)
	}

	if len(result) == 0 {
		return nil, errors.NewNotFoundError("no user rows in result set")
	}
	return result, nil
}

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.Username, &user.Firstname, &user.Surname, &user.Email, &user.RoleID, &user.DepartmentID, &user.Salt, &user.Password, &user.Image, &user.Status, &user.DateCreated); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Firstname, user.Surname, user.Email, user.RoleID, user.DepartmentID, user.Salt, user.Password, user.Image, user.ID)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		logger.Error("error when trying to delete user", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.DbConn().Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find user by status statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Firstname, &user.Surname, &user.Email, &user.RoleID, &user.DepartmentID, &user.Image, &user.Status, &user.DateCreated); err != nil {
			logger.Error("error when trying to scan user row into user struct", err)
			return nil, mysql_utils.ParseError(err)
		}
		result = append(result, *user)
	}

	if len(result) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return result, nil
}

func (user *User) FindByUsernameAndPassword() *errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryFindByUserAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare find user by username and password statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Username, user.Password)
	if getErr := result.Scan(&user.ID, &user.Username, &user.Firstname, &user.Surname, &user.Email, &user.Image, &user.Status, &user.DateCreated); getErr != nil {
		if strings.Contains(getErr.Error(), "no rows in result set") {
			return errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by username and password", getErr)
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) GetSalt() *errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryGetSalt)
	if err != nil {
		logger.Error("error when trying to prepare user salt by username statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	row := stmt.QueryRow(user.Username)
	if getErr := row.Scan(&user.Salt); getErr != nil {
		if strings.Contains(getErr.Error(), "no rows in result set") {
			return errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user salt by username", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}
