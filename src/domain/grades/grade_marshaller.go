package grades

type PublicGrade struct {
	ID int64 `json:"id"`
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
			ID: grade.ID,
		}
	}

	return grade
}
