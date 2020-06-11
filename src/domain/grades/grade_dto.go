package grades

import "github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"

type Grade struct {
	ID        int64 `json:"id"`
	UserID    int64 `json:"user_id"`
	SectionID int64 `json:"section_id"`
	Grade     int64 `json:"grade"`
}

type Grades []Grade

func (grade Grade) Validate() rest_errors.RestErr {
	if grade.UserID <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}

	if grade.SectionID <= 0 {
		return rest_errors.NewBadRequestError("invalid section id")
	}

	if grade.Grade < 0 {
		return rest_errors.NewBadRequestError("invalid grade user")
	}

	return nil
}
