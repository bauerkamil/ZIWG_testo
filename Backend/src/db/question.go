package db

import (
	"github.com/gofrs/uuid"
	"src/model"
)

func AddQuestionToDB(question *model.Question) error {
	dbC := DB
	result := dbC.Create(question)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetQuestionsFromDB() ([]model.Question, error) {
	dbC := DB
	var questions []model.Question
	result := dbC.Find(&questions)
	if result.Error != nil {
		return nil, result.Error
	}
	return questions, nil
}

func GetQuestionFromDB(id uuid.UUID) (*model.Question, error) {
	dbC := DB
	var question model.Question
	result := dbC.First(&question, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &question, nil
}

func UpdateQuestionInDB(newQuestion *model.Question) error {
	dbC := DB
	var Question model.Question
	result := dbC.First(&Question, newQuestion.Id)
	if result.Error != nil {
		return result.Error
	}
	dbC.Save(newQuestion)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteQuestionFromDB(id uuid.UUID) error {
	dbC := DB
	result := dbC.Delete(&model.Question{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
