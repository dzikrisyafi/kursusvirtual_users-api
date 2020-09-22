package menu

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/menu"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var menu menu.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, err := services.MenuService.CreateMenu(menu)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	res := rest_resp.NewStatusCreated("success created menu", result)
	c.JSON(res.Status(), res)
}

func Get(c *gin.Context) {
	menuID, err := controller_utils.GetIDInt(c.Param("menu_id"), "menu id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	result, err := services.MenuService.GetMenu(menuID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	res := rest_resp.NewStatusOK("success get menu", result)
	c.JSON(res.Status(), res)
}

func GetAll(c *gin.Context) {
	result, err := services.MenuService.GetAllMenu()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	res := rest_resp.NewStatusOK("success get menu", result)
	c.JSON(res.Status(), res)
}

func GetAllByRoleID(c *gin.Context) {
	roleID, err := controller_utils.GetIDInt(c.Param("role_id"), "role id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	result, err := services.MenuService.GetAllMenuByRoleID(roleID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	res := rest_resp.NewStatusOK("success get menu", result)
	c.JSON(res.Status(), res)

}

func Update(c *gin.Context) {
	menuID, err := controller_utils.GetIDInt(c.Param("menu_id"), "menu id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var menu menu.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	menu.ID = menuID
	result, saveErr := services.MenuService.UpdateMenu(menu)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	res := rest_resp.NewStatusOK("success updated menu", result)
	c.JSON(res.Status(), res)
}

func Delete(c *gin.Context) {
	menuID, err := controller_utils.GetIDInt(c.Param("menu_id"), "menu id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := services.MenuService.DeleteMenu(menuID); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleted menu", "status": http.StatusOK})
}
