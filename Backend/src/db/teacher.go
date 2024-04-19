package db

import "github.com/gofrs/uuid"

type Teacher struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	SecondName string    `json:"second_name"`
	Surname    string    `json:"surname"`
}

func AddTeacherToDB(teacher *Teacher) error {
	dbC := DB
	result := dbC.Create(teacher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetTeachersFromDB() ([]Teacher, error) {
	dbC := DB
	var teachers []Teacher
	result := dbC.Find(&teachers)
	if result.Error != nil {
		return nil, result.Error
	}
	return teachers, nil
}

//make teacher optional

func GetTeacherFromDB(id uuid.UUID) (*Teacher, error) {
	dbC := DB
	var teacher Teacher
	result := dbC.First(&teacher, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &teacher, nil
}

func UpdateTeacherInDB(newTeacher *Teacher) error {
	dbC := DB
	var teacher Teacher
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
	result := dbC.Delete(&Teacher{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
