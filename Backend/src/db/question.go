package db

import "github.com/gofrs/uuid"

type Question struct {
	Id      uuid.UUID `json:"id"`
	Body    string    `json:"body"`
	ImgFile string    `json:"img_file"`
	TestId  uuid.UUID `json:"test_id"`
}

func AddQuestionToDB(question *Question) error {
	dbC := DB
	result := dbC.Create(question)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetQuestionsFromDB() ([]Question, error) {
	dbC := DB
	var questions []Question
	result := dbC.Find(&questions)
	if result.Error != nil {
		return nil, result.Error
	}
	return questions, nil
}

func GetQuestionFromDB(id uuid.UUID) (*Question, error) {
	dbC := DB
	var question Question
	result := dbC.First(&question, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &question, nil
}

func UpdateQuestionInDB(newQuestion *Question) error {
	dbC := DB
	var Question Question
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
	result := dbC.Delete(&Question{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
