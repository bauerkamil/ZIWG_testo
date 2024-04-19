package db

import (
	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Test struct {
	Id        uuid.UUID             `json:"id"`
	Name      string                `json:"name"`
	CreatedBy string                `json:"created_by"`
	ChangedBy string                `json:"changed_by"`
	CreatedAt timestamppb.Timestamp `json:"created_at"`
	CourseId  uuid.UUID             `json:"course_id"`
}

func AddTestToDB(test *Test) error {
	dbC := DB
	result := dbC.Create(test)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetTestsFromDB() ([]Test, error) {
	dbC := DB
	var tests []Test
	result := dbC.Find(&tests)
	if result.Error != nil {
		return nil, result.Error
	}
	return tests, nil
}

func GetTestFromDB(id uuid.UUID) (*Test, error) {
	dbC := DB
	var test Test
	result := dbC.First(&test, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &test, nil
}

func UpdateTestInDB(newTest *Test) error {
	dbC := DB
	var test Test
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
	if !changed {
		test.ChangedBy = newTest.ChangedBy
		result = dbC.Save(&test)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func DeleteTestFromDB(id uuid.UUID) error {
	dbC := DB
	result := dbC.Delete(&Test{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
