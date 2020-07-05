package grades

import "github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"

type Grade struct {
	ID         int `json:"id"`
	UserID     int `json:"user_id"`
	ActivityID int `json:"activity_id"`
	CourseID   int `json:"course_id"`
	Grade      int `json:"grade"`
}

type Grades []Grade

func (grade Grade) Validate() rest_errors.RestErr {
	if grade.UserID <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}

	if grade.ActivityID <= 0 {
		return rest_errors.NewBadRequestError("invalid section id")
	}

	if grade.CourseID < 0 {
		return rest_errors.NewBadRequestError("invalid course id")
	}

	if grade.Grade < 0 {
		return rest_errors.NewBadRequestError("invalid grade user")
	}

	return nil
}
