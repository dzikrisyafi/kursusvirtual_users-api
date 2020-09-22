package submenu

import (
	"errors"

	"github.com/dzikrisyafi/kursusvirtual_users-api/src/datasources/mysql/users_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

const (
	queryInsertSubMenu = `INSERT INTO sub_menu (menu_id, title, url, icon, is_active) VALUES (?, ?, ?, ?, ?);`
	queryGetSubMenu    = `SELECT menu_id, title, url, icon, is_active FROM sub_menu WHERE id=?;`
	queryGetAllSubMenu = `SELECT id, menu_id, title, url, icon, is_active FROM sub_menu;`
	queryUpdateSubMenu = `UPDATE sub_menu SET menu_id=?, title=?, url=?, icon=?, is_active=? WHERE id=?;`
	queryDeleteSubMenu = `DELETE FROM sub_menu WHERE id=?;`
)

func (submenu *SubMenu) Save(isActive int) rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryInsertSubMenu)
	if err != nil {
		logger.Error("error when trying to prepare save submenu statement", err)
		return rest_errors.NewInternalServerError("error when trying to save submenu", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(submenu.MenuID, submenu.Title, submenu.URL, submenu.Icon, isActive)
	if saveErr != nil {
		logger.Error("error when trying to save submenu", err)
		return rest_errors.NewInternalServerError("error when trying to save submenu", errors.New("database error"))
	}

	submenuID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to ge last insert id after creating a new submenu", err)
		return rest_errors.NewInternalServerError("error when trying to save submenu", err)
	}
	submenu.ID = int(submenuID)

	return nil
}

func (submenu *SubMenu) Get() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryGetSubMenu)
	if err != nil {
		logger.Error("error when trying to prepare get submenu by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to get submenu", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(submenu.ID)
	var isActive int
	if getErr := result.Scan(&submenu.MenuID, &submenu.Title, &submenu.URL, &submenu.Icon, &isActive); getErr != nil {
		logger.Error("error when trying to get submenu by id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get submenu", errors.New("database error"))
	}

	if isActive == 1 {
		submenu.IsActive = true
	} else {
		submenu.IsActive = false
	}

	return nil
}

func (submenu *SubMenu) GetAll() ([]SubMenu, rest_errors.RestErr) {
	stmt, err := users_db.DbConn().Prepare(queryGetAllSubMenu)
	if err != nil {
		logger.Error("error when trying to prepare get all submenu statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get submenu", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		logger.Error("error when trying to get submenu", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get submenu", errors.New("database error"))
	}
	defer rows.Close()

	result := make([]SubMenu, 0)
	var isActive int
	for rows.Next() {
		if err := rows.Scan(&submenu.ID, &submenu.MenuID, &submenu.Title, &submenu.URL, &submenu.Icon, &isActive); err != nil {
			logger.Error("error when trying to scan submenu rows into submenu struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get submenu", errors.New("database error"))
		}

		if isActive == 1 {
			submenu.IsActive = true
		} else {
			submenu.IsActive = false
		}

		result = append(result, *submenu)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError("no submenu rows in result set")
	}

	return result, nil
}

func (submenu *SubMenu) Update(isActive int) rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryUpdateSubMenu)
	if err != nil {
		logger.Error("error when trying to prepare update submenu statement", err)
		return rest_errors.NewInternalServerError("error when trying to update submenu", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(submenu.MenuID, submenu.Title, submenu.URL, submenu.Icon, isActive, submenu.ID)
	if err != nil {
		logger.Error("error when trying to update submenu", err)
		return rest_errors.NewInternalServerError("error when trying to update submenu", errors.New("database error"))
	}

	return nil
}

func (submenu *SubMenu) Delete() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryDeleteSubMenu)
	if err != nil {
		logger.Error("error when trying to prepare delete submenu statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete submenu", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(submenu.ID)
	if err != nil {
		logger.Error("error when trying to delete submenu", err)
		return rest_errors.NewInternalServerError("error when trying to delete submenu", errors.New("database error"))
	}

	return nil
}
