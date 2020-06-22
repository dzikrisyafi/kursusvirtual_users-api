package departments

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/departments"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var department departments.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, err := services.DepartmentsService.CreateDepartment(department)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	resp := rest_resp.NewStatusCreated("success created department", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Get(c *gin.Context) {
	departmentID, err := controller_utils.GetIDInt(c.Param("department_id"), "department id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	result, err := services.DepartmentsService.GetDepartment(departmentID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	resp := rest_resp.NewStatusOK("success get department", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func GetAll(c *gin.Context) {
	result, err := services.DepartmentsService.GetAllDepartment()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	resp := rest_resp.NewStatusOK("success get department", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Update(c *gin.Context) {
	departmentID, err := controller_utils.GetIDInt(c.Param("department_id"), "department id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var department departments.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	department.ID = departmentID
	result, saveErr := services.DepartmentsService.UpdateDepartment(department)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	resp := rest_resp.NewStatusOK("success updated department", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Delete(c *gin.Context) {
	departmentID, err := controller_utils.GetIDInt(c.Param("department_id"), "department id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := services.DepartmentsService.DeleteDepartment(departmentID); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleted department", "status": http.StatusOK})
}
