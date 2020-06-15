package app

import "github.com/dzikrisyafi/kursusvirtual_grades-api/src/controllers/grades"

func mapUrls() {
	router.POST("/grades", grades.Create)
	router.GET("/grades/:grade_id", grades.Get)
	router.DELETE("/grades/:grade_id", grades.Delete)
	router.PUT("/grades/:grade_id", grades.Update)

	router.GET("internal/grade/user/:user_id/:activity_id", grades.GetByUserAndActivityID)
	router.GET("internal/grades/users/:user_id/:course_id", grades.GetAllByUserAndCourseID)
	router.DELETE("/internal/grades/users/:user_id", grades.DeleteAllByUserID)
	router.DELETE("/internal/grades/courses/:course_id", grades.DeleteAllByCourseID)
}
