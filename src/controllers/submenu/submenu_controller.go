package submenu

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/submenu"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	submenuID, err := controller_utils.GetIDInt(c.Param("submenu_id"), "submenu id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	result, err := services.SubMenuService.GetSubMenu(submenuID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	res := rest_resp.NewStatusOK("success get submenu", result)
	c.JSON(res.Status(), res)
}

func GetAll(c *gin.Context) {
	result, err := services.SubMenuService.GetAllSubMenu()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	res := rest_resp.NewStatusOK("success get submenu", result)
	c.JSON(res.Status(), res)
}

func Create(c *gin.Context) {
	var submenu submenu.SubMenu
	if err := c.ShouldBindJSON(&submenu); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.SubMenuService.CreateSubMenu(submenu)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	res := rest_resp.NewStatusCreated("succcess created submenu", result)
	c.JSON(res.Status(), res)
}

func Update(c *gin.Context) {
	submenuID, err := controller_utils.GetIDInt(c.Param("submenu_id"), "submenu id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var submenu submenu.SubMenu
	if err := c.ShouldBindJSON(&submenu); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	submenu.ID = submenuID
	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.SubMenuService.UpdateSubMenu(isPartial, submenu)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	res := rest_resp.NewStatusOK("success updated submenu", result)
	c.JSON(res.Status(), res)
}

func Delete(c *gin.Context) {
	submenuID, err := controller_utils.GetIDInt(c.Param("submenu_id"), "submenu id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := services.SubMenuService.DeleteSubMenu(submenuID); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleted submenu", "status": http.StatusOK})
}
