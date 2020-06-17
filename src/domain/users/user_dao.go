package users

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_users-api/src/datasources/mysql/users_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
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

// data access object to make transaction into database users_db.

func (user *User) Save() rest_errors.RestErr {
	// prepare query statement to insert user into database
	stmt, err := users_db.DbConn().Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}
	// close the statement before return function
	defer stmt.Close()

	// execute a prepared query statement with the given arguments
	insertResult, saveErr := stmt.Exec(user.Username, user.Firstname, user.Surname, user.Email, user.RoleID, user.DepartmentID, user.Salt, user.Password, user.Image, user.Status, user.DateCreated)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}

	// generating last insert id after creating a new user
	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}
	user.ID = int(userID)

	return nil
}

func (user *User) GetAllUser() ([]User, rest_errors.RestErr) {
	// prepare query statement to get all user from database
	stmt, err := users_db.DbConn().Prepare(queryGetAllUser)
	if err != nil {
		logger.Error("error when trying to prepare get all user statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get users", errors.New("database error"))
	}
	// close the statement before return function
	defer stmt.Close()

	// execute a prepared query statement to get all user
	rows, getErr := stmt.Query()
	if getErr != nil {
		logger.Error("error when trying to get users", getErr)
		return nil, rest_errors.NewInternalServerError("error when trying to get users", errors.New("database error"))
	}
	// close the rows, preventing futher enumeration
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next() {
		// copies data from database to user struct
		if err := rows.Scan(&user.ID, &user.Username, &user.Firstname, &user.Surname, &user.Email, &user.RoleID, &user.DepartmentID, &user.Image, &user.Status, &user.DateCreated); err != nil {
			logger.Error("error when trying to scan user rows into user struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get users", errors.New("database error"))
		}

		result = append(result, *user)
	}

	// check the length of result, to make sure if the result is not nil
	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError("no user rows in result set")
	}
	return result, nil
}

func (user *User) Get() rest_errors.RestErr {
	// prepare query statement to get user from database
	stmt, err := users_db.DbConn().Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}
	// close the statement before return function
	defer stmt.Close()

	// execute a prepared query statement with the given arguments
	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.Username, &user.Firstname, &user.Surname, &user.Email, &user.RoleID, &user.DepartmentID, &user.Salt, &user.Password, &user.Image, &user.Status, &user.DateCreated); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}

	return nil
}

func (user *User) Update() rest_errors.RestErr {
	// prepare query statement to update user
	stmt, err := users_db.DbConn().Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return rest_errors.NewInternalServerError("error when trying to update user", errors.New("database error"))
	}
	// close the statement before return function
	defer stmt.Close()

	// execute a prepared query statement with the given arguments
	_, err = stmt.Exec(user.Username, user.Firstname, user.Surname, user.Email, user.RoleID, user.DepartmentID, user.Salt, user.Password, user.Image, user.ID)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return rest_errors.NewInternalServerError("error when trying to update user", errors.New("database error"))
	}

	return nil
}

func (user *User) Delete() rest_errors.RestErr {
	// prepare query statement to delete user by id
	stmt, err := users_db.DbConn().Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}
	// close the statement before return function
	defer stmt.Close()

	// execute a prepared query statement with the given arguments
	if _, err = stmt.Exec(user.ID); err != nil {
		logger.Error("error when trying to delete user by id", err)
		return rest_errors.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, rest_errors.RestErr) {
	// prepare query statement to find user by status
	stmt, err := users_db.DbConn().Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get users", errors.New("database error"))
	}
	// close the statement before return function
	defer stmt.Close()

	// execute a prepared query statement with the given arguments
	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find user by status", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get users", errors.New("database error"))
	}
	// close the rows, to prevent futher enumeration
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Firstname, &user.Surname, &user.Email, &user.RoleID, &user.DepartmentID, &user.Image, &user.Status, &user.DateCreated); err != nil {
			logger.Error("error when trying to scan user row into user struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get users", errors.New("database error"))
		}
		result = append(result, *user)
	}

	// check the length of result, to make sure if the result is not nil
	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return result, nil
}

func (user *User) FindByUsernameAndPassword() rest_errors.RestErr {
	// prepare query statement to find user by username and password
	stmt, err := users_db.DbConn().Prepare(queryFindByUserAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare find user by username and password statement", err)
		return rest_errors.NewInternalServerError("error when trying to find user", errors.New("database error"))
	}
	// close the statement before return function
	defer stmt.Close()

	// execute a prepared query statement with the given arguments
	result := stmt.QueryRow(user.Username, user.Password)
	if getErr := result.Scan(&user.ID, &user.Username, &user.Firstname, &user.Surname, &user.Email, &user.Image, &user.Status, &user.DateCreated); getErr != nil {
		if strings.Contains(getErr.Error(), "no rows in result set") {
			return rest_errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by username and password", getErr)
		return rest_errors.NewInternalServerError("error when trying to find user", errors.New("database error"))
	}

	return nil
}

func (user *User) GetSalt() rest_errors.RestErr {
	// prepare query statement to get salt by username
	stmt, err := users_db.DbConn().Prepare(queryGetSalt)
	if err != nil {
		logger.Error("error when trying to prepare user salt by username statement", err)
		return rest_errors.NewInternalServerError("error when trying to get user salt", errors.New("database error"))
	}
	// close the statement before return function
	defer stmt.Close()

	// execute a prepared query statement with the given arguments
	row := stmt.QueryRow(user.Username)
	if getErr := row.Scan(&user.Salt); getErr != nil {
		if strings.Contains(getErr.Error(), "no rows in result set") {
			return rest_errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user salt by username", err)
		return rest_errors.NewInternalServerError("error when trying to get user salt", errors.New("database error"))
	}

	return nil
}
