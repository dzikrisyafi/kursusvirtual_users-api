package menu

import (
	"errors"

	"github.com/dzikrisyafi/kursusvirtual_users-api/src/datasources/mysql/users_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

const (
	queryInsertMenu         = `INSERT INTO menu (name) VALUES (?);`
	queryGetMenu            = `SELECT name FROM menu WHERE id=?;`
	queryGetAllMenu         = `SELECT id, name FROM menu;`
	queryGetMenuByRoleID    = `SELECT menu.id, menu.name FROM menu JOIN access_menu ON menu.id=access_menu.menu_id WHERE access_menu.role_id=? ORDER BY access_menu.menu_id ASC;`
	queryGetSubmenuByMenuID = `SELECT sub_menu.id, menu_id, title, url, icon, is_active FROM sub_menu JOIN menu ON sub_menu.menu_id=menu.id WHERE sub_menu.menu_id=?;`
	queryUpdateMenu         = `UPDATE menu SET name=? WHERE id=?;`
	queryDeleteMenu         = `DELETE FROM menu WHERE id=?;`
)

func (menu *Menu) Save() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryInsertMenu)
	if err != nil {
		logger.Error("error when trying to prepare save menu statement", err)
		return rest_errors.NewInternalServerError("error when trying to save menu", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(menu.Name)
	if saveErr != nil {
		logger.Error("error when trying to save menu", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save menu", errors.New("database error"))
	}

	menuID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new menu", err)
		return rest_errors.NewInternalServerError("error when trying to save menu", errors.New("database error"))
	}
	menu.ID = int(menuID)

	return nil
}

func (menu *Menu) Get() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryGetMenu)
	if err != nil {
		logger.Error("error when trying to prepare get menu by id", err)
		return rest_errors.NewInternalServerError("error when trying to get menu", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(menu.ID)
	if getErr := result.Scan(&menu.Name); getErr != nil {
		logger.Error("error when trying to get menu by id", err)
		return rest_errors.NewInternalServerError("error when trying to get menu", errors.New("database error"))
	}

	return nil
}

func (menu *Menu) GetAll() ([]Menu, rest_errors.RestErr) {
	stmt, err := users_db.DbConn().Prepare(queryGetAllMenu)
	if err != nil {
		logger.Error("error when trying to prepare get all menu statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get menu", errors.New("database error"))
	}
	defer stmt.Close()

	rows, getErr := stmt.Query()
	if getErr != nil {
		logger.Error("error when trying to get menu", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get menu", errors.New("database error"))
	}
	defer rows.Close()

	result := make([]Menu, 0)
	for rows.Next() {
		if err := rows.Scan(&menu.ID, &menu.Name); err != nil {
			logger.Error("error when trying to scan menu rows into menu struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get menu", errors.New("database error"))
		}

		result = append(result, *menu)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError("no menu rows in result set")
	}

	return result, nil
}

func (menu *AccessMenu) GetMenuByRoleID(roleID int) ([]AccessMenu, rest_errors.RestErr) {
	stmt, err := users_db.DbConn().Prepare(queryGetMenuByRoleID)
	if err != nil {
		logger.Error("error when trying to prepare get menu by role id", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get menu", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(roleID)
	if err != nil {
		logger.Error("error when trying to get menu", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get menu", errors.New("database error"))
	}
	defer rows.Close()

	result := make([]AccessMenu, 0)
	for rows.Next() {
		if err := rows.Scan(&menu.ID, &menu.Name); err != nil {
			logger.Error("error when trying to scan menu rows into menu struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get menu", errors.New("database error"))
		}

		result = append(result, *menu)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError("no menu rows in result set")
	}

	return result, nil
}

func (submenu *AccessSubMenu) GetAllSubmenuByMenuID(menu *AccessMenu) rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryGetSubmenuByMenuID)
	if err != nil {
		logger.Error("error when trying to prepare get submenu by menu id statement", err)
		return rest_errors.NewInternalServerError("error when trying to get submenu", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(menu.ID)
	if err != nil {
		logger.Error("error when trying to get submenu by menu id", err)
		return rest_errors.NewInternalServerError("error when trying to get submenu", errors.New("database error"))
	}
	defer rows.Close()

	var isActive int
	for rows.Next() {
		if err := rows.Scan(&submenu.ID, &submenu.MenuID, &submenu.Title, &submenu.URL, &submenu.Icon, &submenu.IsActive); err != nil {
			logger.Error("error when trying to scan row into submenu struct", err)
			return rest_errors.NewInternalServerError("error when trying to get submenu", errors.New("database error"))
		}

		if isActive == 1 {
			submenu.IsActive = true
		} else {
			submenu.IsActive = false
		}

		menu.SubMenu = append(menu.SubMenu, *submenu)
	}

	return nil
}

func (menu *Menu) Update() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryUpdateMenu)
	if err != nil {
		logger.Error("error when trying to prepare update menu statement", err)
		return rest_errors.NewInternalServerError("error when trying to update menu", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(menu.Name, menu.ID)
	if err != nil {
		logger.Error("error when trying to update menu", err)
		return rest_errors.NewInternalServerError("error when trying to update menu", errors.New("database error"))
	}

	return nil
}

func (menu *Menu) Delete() rest_errors.RestErr {
	stmt, err := users_db.DbConn().Prepare(queryDeleteMenu)
	if err != nil {
		logger.Error("error when trying to prepare delete menu statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete menu", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(menu.ID); err != nil {
		logger.Error("error when trying to delete menu", err)
		return rest_errors.NewInternalServerError("error when trying to delete menu", errors.New("database error"))
	}

	return nil
}
