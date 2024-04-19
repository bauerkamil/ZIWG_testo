package db

import (
	"github.com/gofrs/uuid"
	"src/model"
)

func AddTeacherToDB(teacher *model.Teacher) error {
	dbC := DB
	result := dbC.Create(teacher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetTeachersFromDB() ([]model.Teacher, error) {
	dbC := DB
	var teachers []model.Teacher
	result := dbC.Find(&teachers)
	if result.Error != nil {
		return nil, result.Error
	}
	return teachers, nil
}

//make teacher optional

func GetTeacherFromDB(id uuid.UUID) (*model.Teacher, error) {
	dbC := DB
	var teacher model.Teacher
	result := dbC.First(&teacher, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &teacher, nil
}

func UpdateTeacherInDB(newTeacher *model.Teacher) error {
	dbC := DB
	var teacher model.Teacher
	result := dbC.First(&teacher, newTeacher.Id)
	if result.Error != nil {
		return result.Error
	}
	dbC.Save(newTeacher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteTeacherFromDB(id uuid.UUID) error {
	dbC := DB
	result := dbC.Delete(&model.Teacher{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
