package model

import "github.com/gofrs/uuid"

type AnswerRequest struct {
	QuestionId uuid.UUID `json:"question_id"`
	Body       string    `json:"body"`
	Valid      bool      `json:"valid"`
}

type CourseRequest struct {
	Name       string    `json:"name"`
	TeacherId  uuid.UUID `json:"teacher_id"`
	SchoolYear int       `json:"school_year"`
}

type QuestionRequest struct {
	Body    string    `json:"body"`
	ImgFile string    `json:"img_file"`
	TestId  uuid.UUID `json:"test_id"`
}

type TeacherRequest struct {
	Name       string `json:"name"`
	SecondName string `json:"second_name"`
	Surname    string `json:"surname"`
}

type TestRequest struct {
	Name     string    `json:"name"`
	CourseId uuid.UUID `json:"course_id"`
}
