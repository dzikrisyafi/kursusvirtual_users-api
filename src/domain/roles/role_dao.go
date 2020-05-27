package roles

import (
	"errors"

	"github.com/dzikrisyafi/kursusvirtual_users-api/src/datasources/mysql/users_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

const (
	queryGetRole    = `SELECT id, name FROM roles WHERE id=?;`
	queryGetAllRole = `SELECT id, name FROM roles;`
	queryInsertRole = `INSERT INTO roles(name) VALUES(?);`
	queryUpdateRole = `UPDATE roles SET name=? WHERE id=?;`
	queryDeleteRole = `DELETE FROM roles WHERE id=?;`
)

func (role *Role) Get() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryGetRole)
	if err != nil {
		logger.Error("error when trying to prepare get role statement", err)
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(role.ID)
	if getErr := result.Scan(&role.ID, &role.Name); getErr != nil {
		logger.Error("error when trying to get role by id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get user by id", errors.New("database error"))
	}

	return nil
}

func (role *Role) GetAllRole() ([]Role, rest_errors.RestErr) {
	stmt, err := users_db.DbConn().Prepare(queryGetAllRole)
	if err != nil {
		logger.Error("error when trying to prepare get all role statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get roles", errors.New("database error"))
	}
	defer stmt.Close()

	rows, getErr := stmt.Query()
	if getErr != nil {
		logger.Error("error when trying to get roles", getErr)
		return nil, rest_errors.NewInternalServerError("error when trying to get roles", errors.New("database error"))
	}
	defer rows.Close()

	result := make([]Role, 0)
	for rows.Next() {
		if err := rows.Scan(&role.ID, &role.Name); err != nil {
			logger.Error("error when trying to scan role rows into role struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get roles", errors.New("database error"))
		}

		result = append(result, *role)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError("no role rows in result set")
	}
	return result, nil
}

func (role *Role) Save() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryInsertRole)
	if err != nil {
		logger.Error("error when trying to prepare save role statement", err)
		return rest_errors.NewInternalServerError("error when trying to save role", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(role.Name)
	if saveErr != nil {
		logger.Error("error when trying to save role", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save role", errors.New("database error"))
	}

	roleID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new role", err)
		return rest_errors.NewInternalServerError("error when trying to save role", errors.New("database error"))
	}
	role.ID = roleID

	return nil
}

func (role *Role) Delete() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryDeleteRole)
	if err != nil {
		logger.Error("error when trying to prepare delete role by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete role", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(role.ID); err != nil {
		logger.Error("error when trying to delete role by id", err)
		return rest_errors.NewInternalServerError("error when trying to delete role", errors.New("database error"))
	}

	return nil
}

func (role *Role) Update() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryUpdateRole)
	if err != nil {
		logger.Error("error when trying to prepare update role statement", err)
		return rest_errors.NewInternalServerError("error when trying to update role", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(role.Name, role.ID)
	if err != nil {
		logger.Error("error when trying to update role", err)
		return rest_errors.NewInternalServerError("error when trying to update role", errors.New("database error"))
	}

	return nil
}
