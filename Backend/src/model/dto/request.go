package dto

import "github.com/gofrs/uuid"

// Teacher models
type TeacherRequest struct {
	Name       string `json:"name"`
	SecondName string `json:"second_name"`
	Surname    string `json:"surname"`
}

// Course models
type CourseRequest struct {
	Name       string    `json:"name"`
	TeacherId  uuid.UUID `json:"teacher_id"`
	UsosId     string    `json:"usos_id"`
	CourseType string    `json:"course_type"`
}

// Test models
type TestRequest struct {
	Name       string        `json:"name"`
	CourseId   uuid.UUID     `json:"course_id"`
	SchoolYear string        `json:"school_year"`
	Questions  []SubQuestion `json:"questions"`
}

type EditTestRequest struct {
	Name       string    `json:"name"`
	CourseId   uuid.UUID `json:"course_id"`
	SchoolYear string    `json:"school_year"`
}

// Question models
type QuestionRequest struct {
	Body    string      `json:"body"`
	ImgFile string      `json:"img_file"`
	TestId  uuid.UUID   `json:"test_id"`
	Answers []SubAnswer `json:"answers"`
}

type EditQuestionRequest struct {
	Body    string          `json:"body"`
	ImgFile string          `json:"img_file"`
	Answers []EditSubAnswer `json:"answers"`
}

type SubQuestion struct {
	Body    string      `json:"body"`
	ImgFile string      `json:"img_file"`
	Answers []SubAnswer `json:"answers"`
}

// Answer models
type AnswerRequest struct {
	QuestionId uuid.UUID `json:"question_id"`
	Body       string    `json:"body"`
	Valid      bool      `json:"valid"`
}

type SubAnswer struct {
	Body  string `json:"body"`
	Valid bool   `json:"valid"`
}

type EditSubAnswer struct {
	SubAnswer
	Id *uuid.UUID `json:"id"`
}
