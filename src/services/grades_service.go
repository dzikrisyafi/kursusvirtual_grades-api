package services

import (
	"github.com/dzikrisyafi/kursusvirtual_grades-api/src/domain/grades"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	GradesService gradesServiceInterface = &gradesService{}
)

type gradesService struct{}

type gradesServiceInterface interface {
	CreateGrade(grades.Grade) (*grades.Grade, rest_errors.RestErr)
	GetGrade(int, int) (*grades.Grade, rest_errors.RestErr)
	GetGradeByID(int) (*grades.Grade, rest_errors.RestErr)
	GetAll(int) (grades.Grades, rest_errors.RestErr)
	UpdateGrade(grades.Grade) (*grades.Grade, rest_errors.RestErr)
	DeleteGrade(int) rest_errors.RestErr
	DeleteUserGradeByUserID(int) rest_errors.RestErr
	DeleteUserGradeByCourseID(int) rest_errors.RestErr
}

func (s *gradesService) CreateGrade(grade grades.Grade) (*grades.Grade, rest_errors.RestErr) {
	if err := grade.Validate(); err != nil {
		return nil, err
	}

	if err := grade.Save(); err != nil {
		return nil, err
	}

	return &grade, nil
}

func (s *gradesService) GetGrade(userID int, activityID int) (*grades.Grade, rest_errors.RestErr) {
	result := &grades.Grade{UserID: userID, ActivityID: activityID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *gradesService) GetGradeByID(gradeID int) (*grades.Grade, rest_errors.RestErr) {
	result := &grades.Grade{ID: gradeID}
	if err := result.GetByID(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *gradesService) GetAll(userID int) (grades.Grades, rest_errors.RestErr) {
	dao := &grades.Grade{UserID: userID}
	return dao.GetAll()
}

func (s *gradesService) UpdateGrade(grade grades.Grade) (*grades.Grade, rest_errors.RestErr) {
	current, err := s.GetGradeByID(grade.ID)
	if err != nil {
		return nil, err
	}

	if err := grade.Validate(); err != nil {
		current.Grade = grade.Grade
	}

	return current, nil
}

func (s *gradesService) DeleteGrade(gradeID int) rest_errors.RestErr {
	dao := &grades.Grade{ID: gradeID}
	return dao.Delete()
}

func (s *gradesService) DeleteUserGradeByUserID(userID int) rest_errors.RestErr {
	dao := &grades.Grade{UserID: userID}
	return dao.DeleteUserGradeByUserID()
}

func (s *gradesService) DeleteUserGradeByCourseID(courseID int) rest_errors.RestErr {
	dao := &grades.Grade{CourseID: courseID}
	return dao.DeleteUserGradeByCourseID()
}
