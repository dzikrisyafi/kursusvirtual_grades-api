package app

import "github.com/dzikrisyafi/kursusvirtual_grades-api/src/controllers/grades"

func mapUrls() {
	router.POST("/grades", grades.Create)
	router.GET("/grades/:user_id/:activity_id", grades.Get)
	router.GET("/grades/:user_id", grades.GetAll)
	router.DELETE("/grades/:grade_id", grades.Delete)
	router.PUT("/grades/:grade_id", grades.Update)

	router.DELETE("/internal/grades/users/:user_id", grades.DeleteAllByUserID)
	router.DELETE("/internal/grades/courses/:course_id", grades.DeleteAllByCourseID)
}
