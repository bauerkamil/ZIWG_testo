package dto

import "github.com/gofrs/uuid"

type AnswerRequest struct {
	QuestionId uuid.UUID `json:"question_id"`
	Body       string    `json:"body"`
	Valid      bool      `json:"valid"`
}

type CourseRequest struct {
	Name       string    `json:"name"`
	TeacherId  uuid.UUID `json:"teacher_id"`
	UsosId     string    `json:"usos_id"`
	CourseType string    `json:"course_type"`
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
	Name       string    `json:"name"`
	CourseId   uuid.UUID `json:"course_id"`
	SchoolYear string    `json:"school_year"`
}
