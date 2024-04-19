package db

import (
	"github.com/gofrs/uuid"
	"src/model"
)

func AddAnswerToDB(answer *model.Answer) error {
	dbC := DB
	result := dbC.Create(answer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAnswersFromDB() ([]model.Answer, error) {
	dbC := DB
	var answers []model.Answer
	result := dbC.Find(&answers)
	if result.Error != nil {
		return nil, result.Error
	}
	return answers, nil
}

func GetAnswerFromDB(id uuid.UUID) (*model.Answer, error) {
	dbC := DB
	var answer model.Answer
	result := dbC.First(&answer, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &answer, nil
}

func UpdateAnswerInDB(newAnswer *model.Answer) error {
	dbC := DB
	var answer model.Answer
	result := dbC.First(&answer, newAnswer.Id)
	if result.Error != nil {
		return result.Error
	}
	dbC.Save(newAnswer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteAnswerFromDB(id uuid.UUID) error {
	dbC := DB
	result := dbC.Delete(&model.Answer{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
