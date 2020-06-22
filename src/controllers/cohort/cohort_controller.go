package cohort

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/domain/cohort"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var cohort cohort.Cohort
	if err := c.ShouldBindJSON(&cohort); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, err := services.CohortService.CreateCohort(cohort)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	cohortID, err := controller_utils.GetIDInt(c.Param("cohort_id"), "cohort id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	result, err := services.CohortService.GetCohort(cohortID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	resp := rest_resp.NewStatusOK("success get cohort", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(http.StatusOK, resp)
}

func GetAll(c *gin.Context) {
	result, err := services.CohortService.GetAllCohort()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	resp := rest_resp.NewStatusOK("success get cohort", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(http.StatusOK, resp)
}

func Update(c *gin.Context) {
	cohortID, err := controller_utils.GetIDInt(c.Param("cohort_id"), "cohort id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var cohort cohort.Cohort
	if err := c.ShouldBindJSON(&cohort); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	cohort.ID = cohortID
	result, err := services.CohortService.UpdateCohort(cohort)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	cohortID, err := controller_utils.GetIDInt(c.Param("cohort_id"), "cohort id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := services.CohortService.DeleteCohort(cohortID); err != nil {
		c.JSON(err.Status(), err)
		return
	}
}
