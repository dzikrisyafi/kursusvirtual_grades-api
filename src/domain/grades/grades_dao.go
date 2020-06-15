package grades

import (
	"errors"

	"github.com/dzikrisyafi/kursusvirtual_grades-api/src/datasources/mysql/grades_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

const (
	queryInsertGrade           = `INSERT INTO grades(user_id, activity_id, grade, course_id) VALUES(?, ?, ?, ?);`
	queryGetGrade              = `SELECT id, grade, course_id FROM grades WHERE user_id=? AND activity_id=?;`
	queryGetGradeByID          = `SELECT user_id, activity_id, grade, course_id FROM grades WHERE id=?;`
	queryGetAllUserGrade       = `SELECT id, activity_id, grade, course_id FROM grades WHERE user_id=? AND course_id=?;`
	queryUpdateGrade           = `UPDATE grades SET grade=? WHERE id=?;`
	queryDeleteGrade           = `DELETE FROM grades WHERE id=?;`
	queryDeleteGradeByUserID   = `DELETE FROM grades WHERE user_id=?;`
	queryDeleteGradeByCourseID = `DELETE FROM grades WHERE course_id=?;`
)

func (grade *Grade) Save() rest_errors.RestErr {
	stmt, err := grades_db.DbConn().Prepare(queryInsertGrade)
	if err != nil {
		logger.Error("error when trying to prepare save grade user statement", err)
		return rest_errors.NewInternalServerError("error when trying to save grade user", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(grade.UserID, grade.ActivityID, grade.Grade, grade.CourseID)
	if saveErr != nil {
		logger.Error("error when trying to save grade user", err)
		return rest_errors.NewInternalServerError("error when trying to save grade user", errors.New("database error"))
	}

	gradeID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new grade", err)
		return rest_errors.NewInternalServerError("error when trying to save grade user", errors.New("database error"))
	}
	grade.ID = int(gradeID)

	return nil
}

func (grade *Grade) GetByUserAndActivityID() rest_errors.RestErr {
	stmt, err := grades_db.DbConn().Prepare(queryGetGrade)
	if err != nil {
		logger.Error("error when trying to prepare get user grade by user and section id statement", err)
		return rest_errors.NewInternalServerError("error when trying to save user grade", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(grade.UserID, grade.ActivityID)
	if getErr := result.Scan(&grade.ID, &grade.Grade, &grade.CourseID); getErr != nil {
		logger.Error("error when trying to get user grade", getErr)
		return rest_errors.NewInternalServerError("error when trying to get user grade", errors.New("database error"))
	}

	return nil
}

func (grade *Grade) GetByUserAndCourseID() rest_errors.RestErr {
	stmt, err := grades_db.DbConn().Prepare(queryGetGradeByID)
	if err != nil {
		logger.Error("error when trying to prepare get user grade by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to save user grade", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(&grade.ID)
	if getErr := result.Scan(&grade.UserID, &grade.ActivityID, &grade.Grade, &grade.CourseID); getErr != nil {
		logger.Error("error when trying to get user grade by id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get user grade", errors.New("database error"))
	}

	return nil
}

func (grade *Grade) GetAll() ([]Grade, rest_errors.RestErr) {
	stmt, err := grades_db.DbConn().Prepare(queryGetAllUserGrade)
	if err != nil {
		logger.Error("error when trying to prepare get all user grade statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get all user grade", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(grade.UserID, grade.CourseID)
	if err != nil {
		logger.Error("error when trying to get all user grade", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get all user grade", errors.New("database error"))
	}
	defer rows.Close()

	result := make([]Grade, 0)
	for rows.Next() {
		if getErr := rows.Scan(&grade.ID, &grade.ActivityID, &grade.Grade, &grade.CourseID); getErr != nil {
			logger.Error("error when trying to scan user grade rows into user grade struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get all user grade", errors.New("database error"))
		}

		result = append(result, *grade)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError("no user grade in result set")
	}

	return result, nil
}

func (grade *Grade) Update() rest_errors.RestErr {
	stmt, err := grades_db.DbConn().Prepare(queryUpdateGrade)
	if err != nil {
		logger.Error("error when trying to update user grade by id", err)
		return rest_errors.NewInternalServerError("error when trying to update user grade", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(grade.Grade, grade.ID); err != nil {
		logger.Error("error when trying to update user grade by id", err)
		return rest_errors.NewInternalServerError("error when trying to update user grade", errors.New("database error"))
	}

	return nil
}

func (grade *Grade) Delete() rest_errors.RestErr {
	stmt, err := grades_db.DbConn().Prepare(queryDeleteGrade)
	if err != nil {
		logger.Error("error when trying to delete user grade by id", err)
		return rest_errors.NewInternalServerError("error when trying to delete user grade", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(grade.ID); err != nil {
		logger.Error("error when trying to delete user grade by id", err)
		return rest_errors.NewInternalServerError("error when trying to delete user grade", errors.New("database error"))
	}

	return nil
}

func (grade *Grade) DeleteUserGradeByUserID() rest_errors.RestErr {
	stmt, err := grades_db.DbConn().Prepare(queryDeleteGradeByUserID)
	if err != nil {
		logger.Error("error when trying to prepare delete user grade by user id statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete user grade", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(grade.UserID); err != nil {
		logger.Error("error when trying to delete user grade by user id", err)
		return rest_errors.NewInternalServerError("error when trying to delete user grade", errors.New("database error"))
	}

	return nil
}

func (grade *Grade) DeleteUserGradeByCourseID() rest_errors.RestErr {
	stmt, err := grades_db.DbConn().Prepare(queryDeleteGradeByCourseID)
	if err != nil {
		logger.Error("error when trying to prepare delete user grade by course id statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete user grade", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err := stmt.Exec(grade.CourseID); err != nil {
		logger.Error("error when trying to delete user grade by course id", err)
		return rest_errors.NewInternalServerError("error when trying to delete user grade", errors.New("database error"))
	}

	return nil
}
