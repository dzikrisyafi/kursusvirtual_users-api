package departments

import (
	"errors"

	"github.com/dzikrisyafi/kursusvirtual_users-api/src/datasources/mysql/users_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

const (
	queryInsertDepartment = `INSERT INTO departments(name) VALUES(?);`
	queryGetDepartment    = `SELECT name FROM departments WHERE id=?;`
	queryGetAllDepartment = `SELECT id, name FROM departments;`
	queryUpdateDepartment = `UPDATE departments SET name=? WHERE id=?;`
	queryDeleteDepartment = `DELETE FROM departments WHERE id=?;`
)

func (department *Department) Save() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryInsertDepartment)
	if err != nil {
		logger.Error("error when trying to prepare save department statement", err)
		return rest_errors.NewInternalServerError("error when trying to save department", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(department.Name)
	if saveErr != nil {
		logger.Error("error when trying to save department", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save department", errors.New("database error"))
	}

	departmentID, err := insertResult.LastInsertId()
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to get last insert id after creating a new department", errors.New("database error"))
	}
	department.ID = int(departmentID)

	return nil
}

func (department *Department) Get() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryGetDepartment)
	if err != nil {
		logger.Error("error when trying to prepare get department by id", err)
		return rest_errors.NewInternalServerError("error when trying to get department", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(department.ID)
	if getErr := result.Scan(&department.Name); getErr != nil {
		logger.Error("error when trying to get department", err)
		return rest_errors.NewInternalServerError("error when trying to get department", errors.New("database error"))
	}

	return nil
}

func (department *Department) GetAll() ([]Department, rest_errors.RestErr) {
	stmt, err := users_db.DbConn().Prepare(queryGetAllDepartment)
	if err != nil {
		logger.Error("error when trying to prepare get department", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get department", errors.New("database error"))
	}
	defer stmt.Close()

	rows, getErr := stmt.Query()
	if getErr != nil {
		logger.Error("error when trying to get department", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get department", errors.New("database error"))
	}
	defer rows.Close()

	result := make([]Department, 0)
	for rows.Next() {
		if err := rows.Scan(&department.ID, &department.Name); err != nil {
			logger.Error("error when trying to scan department rows to department struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get department", err)
		}

		result = append(result, *department)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError("no department rows in result set")
	}

	return result, nil
}

func (department *Department) Update() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryUpdateDepartment)
	if err != nil {
		logger.Error("error when trying to prepare update department by id", err)
		return rest_errors.NewInternalServerError("error when trying to update department", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err := stmt.Exec(department.Name, department.ID); err != nil {
		logger.Error("error when trying to update department", err)
		return rest_errors.NewInternalServerError("error when trying to update department", errors.New("database error"))
	}

	return nil
}

func (department *Department) Delete() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryDeleteDepartment)
	if err != nil {
		logger.Error("error when trying to prepare delete department by id", err)
		return rest_errors.NewInternalServerError("error when trying to delete department", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err := stmt.Exec(department.ID); err != nil {
		logger.Error("error when trying to delete department", err)
		return rest_errors.NewInternalServerError("error when trying to delete department", errors.New("database error"))
	}

	return nil
}
