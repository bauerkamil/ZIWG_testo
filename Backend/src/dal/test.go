package dal

import (
	"github.com/gofrs/uuid"
	"src/model"
)

func AddTestToDB(test *model.Test) error {
	dbC := DB
	result := dbC.Create(test)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetTestsFromDB() ([]model.Test, error) {
	dbC := DB
	var tests []model.Test
	result := dbC.Preload("Course").
		Preload("Course.Teacher").
		Find(&tests)
	if result.Error != nil {
		return nil, result.Error
	}
	return tests, nil
}

func GetTestFromDB(id uuid.UUID) (*model.Test, error) {
	dbC := DB
	var test model.Test
	result := dbC.Preload("Questions").
		Preload("Questions.Answers").
		First(&test, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &test, nil
}

func UpdateTestInDB(newTest *model.Test) error {
	dbC := DB
	var test model.Test
	result := dbC.First(&test, newTest.Id)
	if result.Error != nil {
		return result.Error
	}
	var changed = false
	if newTest.Name != test.Name {
		test.Name = newTest.Name
		changed = true
	}
	if newTest.CourseId != test.CourseId {
		test.CourseId = newTest.CourseId
		changed = true
	}
	if changed {
		test.ChangedBy = newTest.ChangedBy
		test.ChangedAt = newTest.ChangedAt
		result = dbC.Save(&test)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func DeleteTestFromDB(id uuid.UUID) error {
	dbC := DB
	result := dbC.Delete(&model.Test{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
