package db

import "github.com/gofrs/uuid"

type Answer struct {
	Id         uuid.UUID `json:"id"`
	QuestionId uuid.UUID `json:"question_id"`
	Body       string    `json:"body"`
	Valid      bool      `json:"valid"`
}

func AddAnswerToDB(answer *Answer) error {
	dbC := DB
	result := dbC.Create(answer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAnswersFromDB() ([]Answer, error) {
	dbC := DB
	var answers []Answer
	result := dbC.Find(&answers)
	if result.Error != nil {
		return nil, result.Error
	}
	return answers, nil
}

func GetAnswerFromDB(id uuid.UUID) (*Answer, error) {
	dbC := DB
	var answer Answer
	result := dbC.First(&answer, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &answer, nil
}

func UpdateAnswerInDB(newAnswer *Answer) error {
	dbC := DB
	var answer Answer
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
	result := dbC.Delete(&Answer{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
