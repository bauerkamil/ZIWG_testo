package dal

import (
	"github.com/gofrs/uuid"
	"src/model"
)

func AddCourseToDB(course *model.Course) error {
	dbC := DB
	result := dbC.Create(course)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCoursesFromDB() ([]model.Course, error) {
	dbC := DB
	var courses []model.Course
	result := dbC.Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func GetCourseFromDB(id uuid.UUID) (*model.Course, error) {
	dbC := DB
	var course model.Course
	result := dbC.First(&course, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &course, nil
}

func UpdateCourseInDB(newCourse *model.Course) error {
	dbC := DB
	var course model.Course
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
	result := dbC.Delete(&model.Course{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
