package grades

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_grades-api/src/domain/grades"
	"github.com/dzikrisyafi/kursusvirtual_grades-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var grade grades.Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.GradesService.CreateGrade(grade)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(oauth.IsPublic(c.Request)))
}

func Get(c *gin.Context) {
	userID, idErr := controller_utils.GetID(c.Param("user_id"), "user id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	sectionID, idErr := controller_utils.GetID(c.Param("section_id"), "section id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	grade, getErr := services.GradesService.GetGrade(userID, sectionID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, grade.Marshall(oauth.IsPublic(c.Request)))
}

func GetAll(c *gin.Context) {
	userID, err := controller_utils.GetID(c.Param("user_id"), "user id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	result, getErr := services.GradesService.GetAll(userID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(oauth.IsPublic(c.Request)))
}

func Update(c *gin.Context) {
	gradeID, err := controller_utils.GetID(c.Param("grade_id"), "grade id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var grade grades.Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	grade.ID = gradeID
	result, saveErr := services.GradesService.UpdateGrade(grade)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(oauth.IsPublic(c.Request)))
}

func Delete(c *gin.Context) {
	gradeID, err := controller_utils.GetID(c.Param("grade_id"), "grade id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := services.GradesService.DeleteGrade(gradeID); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleted user grade", "status": http.StatusOK})
}
