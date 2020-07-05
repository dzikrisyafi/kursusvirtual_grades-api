package grades

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_grades-api/src/domain/grades"
	"github.com/dzikrisyafi/kursusvirtual_grades-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
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

	resp := rest_resp.NewStatusCreated("success save user grade", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Get(c *gin.Context) {
	gradeID, idErr := controller_utils.GetIDInt(c.Param("grade_id"), "grade id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	grade, getErr := services.GradesService.GetGradeByID(gradeID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success get user grade", grade.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func GetByUserAndActivityID(c *gin.Context) {
	userID, idErr := controller_utils.GetIDInt(c.Param("user_id"), "user id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	activityID, idErr := controller_utils.GetIDInt(c.Param("activity_id"), "activity id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	grade, getErr := services.GradesService.GetGrade(userID, activityID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success get user grade", grade.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func GetAllByUserAndCourseID(c *gin.Context) {
	userID, err := controller_utils.GetIDInt(c.Param("user_id"), "user id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	courseID, err := controller_utils.GetIDInt(c.Param("course_id"), "course id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	result, getErr := services.GradesService.GetAll(userID, courseID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success get user grade", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Update(c *gin.Context) {
	gradeID, err := controller_utils.GetIDInt(c.Param("grade_id"), "grade id")
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

	resp := rest_resp.NewStatusOK("success updating user grade", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Delete(c *gin.Context) {
	gradeID, err := controller_utils.GetIDInt(c.Param("grade_id"), "grade id")
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

func DeleteAllByUserID(c *gin.Context) {
	userID, err := controller_utils.GetIDInt(c.Param("user_id"), "user id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := services.GradesService.DeleteUserGradeByUserID(userID); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleted user grade", "status": http.StatusOK})
}

func DeleteAllByCourseID(c *gin.Context) {
	courseID, err := controller_utils.GetIDInt(c.Param("course_id"), "user id")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := services.GradesService.DeleteUserGradeByCourseID(courseID); err != nil {
		c.JSON(err.Status(), err)
		return
	}
}
