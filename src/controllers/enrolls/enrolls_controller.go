package enrolls

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	courseID, err := controller_utils.GetIDInt(c.Param("course_id"), "course id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	result, getErr := services.EnrollsService.GetUsersByCourseID(courseID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success get user", result)
	c.JSON(resp.Status(), resp)
}

func Create(c *gin.Context) {
	var request enrolls.Enroll
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, err := services.EnrollsService.CreateEnroll(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Update(c *gin.Context) {
	enrollID, err := controller_utils.GetIDInt(c.Param("enroll_id"), "enroll id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var request enrolls.Enroll
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	request.ID = enrollID
	result, err := services.EnrollsService.UpdateEnroll(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	enrollID, err := controller_utils.GetIDInt(c.Param("enroll_id"), "enroll id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := services.EnrollsService.DeleteEnroll(enrollID); err != nil {
		c.JSON(err.Status(), err)
		return
	}
}
