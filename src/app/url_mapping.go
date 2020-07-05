package app

import (
	"github.com/dzikrisyafi/kursusvirtual_grades-api/src/controllers/grades"
	"github.com/dzikrisyafi/kursusvirtual_middleware/middleware"
)

func mapUrls() {
	gradesGroup := router.Group("/grades")
	gradesGroup.Use(middleware.Auth())
	{
		gradesGroup.POST("", grades.Create)
		gradesGroup.GET("/:grade_id", grades.Get)
		gradesGroup.DELETE("/:grade_id", grades.Delete)
		gradesGroup.PUT("/:grade_id", grades.Update)
	}

	internalGroup := router.Group("/internal")
	internalGroup.Use(middleware.Auth())
	{
		internalGroup.GET("/grade/user/:user_id/:activity_id", grades.GetByUserAndActivityID)
		internalGroup.GET("/grades/users/:user_id/:course_id", grades.GetAllByUserAndCourseID)
		internalGroup.DELETE("/grades/users/:user_id", grades.DeleteAllByUserID)
		internalGroup.DELETE("/grades/courses/:course_id", grades.DeleteAllByCourseID)
	}
}
