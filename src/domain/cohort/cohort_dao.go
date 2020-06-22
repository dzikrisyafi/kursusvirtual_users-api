package cohort

import (
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_users-api/src/datasources/mysql/users_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/juju/errors"
)

const (
	queryInsertCohort = `INSERT INTO cohort(name) VALUES(?);`
	queryGetCohort    = `SELECT name FROM cohort WHERE id=?;`
	queryGetAllCohort = `SELECT id, name FROM cohort;`
	queryUpdateCohort = `UPDATE cohort SET name=? WHERE id=?;`
	queryDeleteCohort = `DELETE FROM cohort WHERE id=?;`
)

func (cohort *Cohort) Save() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryInsertCohort)
	if err != nil {
		logger.Error("error when trying to prepare save cohort statement", err)
		return rest_errors.NewInternalServerError("error when trying to save cohort", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(cohort.Name)
	if saveErr != nil {
		logger.Error("error when trying to save cohort", err)
		return rest_errors.NewInternalServerError("error when trying to save cohort", saveErr)
	}

	cohortID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new cohort", errors.New("database error"))
	}
	cohort.ID = int(cohortID)

	return nil
}

func (cohort *Cohort) Get() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryGetCohort)
	if err != nil {
		logger.Error("error when trying to prepare get cohort by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to get cohort", errors.New("database error"))
	}
	defer stmt.Close()

	row := stmt.QueryRow(cohort.ID)
	if getErr := row.Scan(&cohort.Name); getErr != nil {
		if strings.Contains(getErr.Error(), "no rows") {
			return rest_errors.NewNotFoundError("no cohort matching given id")
		}
		logger.Error("error when trying to get cohort by id", err)
		return rest_errors.NewInternalServerError("error when trying to get cohort", errors.New("database error"))
	}

	return nil
}

func (cohort *Cohort) GetAll() ([]Cohort, rest_errors.RestErr) {
	stmt, err := users_db.DbConn().Prepare(queryGetAllCohort)
	if err != nil {
		logger.Error("error when trying to prepare get cohort statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get cohort", errors.New("database error"))
	}
	defer stmt.Close()

	rows, getErr := stmt.Query()
	if getErr != nil {
		logger.Error("error when trying to get cohort", getErr)
		return nil, rest_errors.NewInternalServerError("error when trying to get cohort", errors.New("database error"))
	}
	defer rows.Close()

	result := make([]Cohort, 0)
	for rows.Next() {
		if err := rows.Scan(&cohort.ID, &cohort.Name); err != nil {
			logger.Error("error when trying to scan cohort rows into cohort struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get cohort", errors.New("database error"))
		}

		result = append(result, *cohort)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError("no cohort rows in result set")
	}

	return result, nil
}

func (cohort *Cohort) Update() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryUpdateCohort)
	if err != nil {
		logger.Error("error when trying to prepare update cohort statement", err)
		return rest_errors.NewInternalServerError("error when trying to update cohort", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err := stmt.Exec(cohort.Name, cohort.ID); err != nil {
		logger.Error("error when trying to update cohort", err)
		return rest_errors.NewInternalServerError("error when trying to update cohort", errors.New("database error"))
	}

	return nil
}

func (cohort *Cohort) Delete() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryDeleteCohort)
	if err != nil {
		logger.Error("error when trying to prepare delete cohort by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete cohort", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err := stmt.Exec(cohort.ID); err != nil {
		logger.Error("error when trying to delete cohort by id", err)
		return rest_errors.NewInternalServerError("error when trying to delete cohort", errors.New("database error"))
	}

	return nil
}
