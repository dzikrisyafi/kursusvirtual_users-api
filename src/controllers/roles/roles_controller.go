package roles

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/roles"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	roleID, err := controller_utils.GetIDInt(c.Param("role_id"), "role id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	role, getErr := services.RolesService.GetRole(roleID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success get role", role.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func GetAll(c *gin.Context) {
	roles, err := services.RolesService.GetAllRole()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	resp := rest_resp.NewStatusOK("success get role", roles.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Create(c *gin.Context) {
	var role roles.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.RolesService.CreateRole(role)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	resp := rest_resp.NewStatusCreated("success created role", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Update(c *gin.Context) {
	roleID, idErr := controller_utils.GetIDInt(c.Param("role_id"), "role id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var role roles.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	role.ID = roleID
	result, saveErr := services.RolesService.UpdateRole(role)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	resp := rest_resp.NewStatusOK("success updated role", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Delete(c *gin.Context) {
	roleID, idErr := controller_utils.GetIDInt(c.Param("role_id"), "role id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.RolesService.DeleteRole(roleID); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleted role", "status": http.StatusOK})
}
