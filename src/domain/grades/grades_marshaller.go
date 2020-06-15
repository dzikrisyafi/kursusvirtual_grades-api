package grades

import "encoding/json"

type PublicGrade struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
}

type PrivateGrade struct {
	ID         int `json:"id"`
	UserID     int `json:"user_id"`
	ActivityID int `json:"activity_id"`
	Grade      int `json:"grade"`
}

func (grades Grades) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(grades))
	for index, grade := range grades {
		result[index] = grade.Marshall(isPublic)
	}

	return result
}

func (grade *Grade) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicGrade{
			ID:     grade.ID,
			UserID: grade.UserID,
		}
	}

	gradeJson, _ := json.Marshal(grade)
	var privateGrade PrivateGrade
	json.Unmarshal(gradeJson, &privateGrade)
	return privateGrade
}
