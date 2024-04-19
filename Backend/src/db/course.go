package db

import "github.com/gofrs/uuid"

type Course struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Teacher    uuid.UUID `json:"teacher"`
	SchoolYear int       `json:"school_year"`
}

func AddCourseToDB(course *Course) error {
	dbC := DB
	result := dbC.Create(course)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCoursesFromDB() ([]Course, error) {
	dbC := DB
	var courses []Course
	result := dbC.Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func GetCourseFromDB(id uuid.UUID) (*Course, error) {
	dbC := DB
	var course Course
	result := dbC.First(&course, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &course, nil
}

func UpdateCourseInDB(newCourse *Course) error {
	dbC := DB
	var course Course
	result := dbC.First(&course, newCourse.Id)
	if result.Error != nil {
		return result.Error
	}
	dbC.Save(newCourse)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteCourseFromDB(id uuid.UUID) error {
	dbC := DB
	result := dbC.Delete(&Course{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
