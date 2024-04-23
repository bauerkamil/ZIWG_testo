package model

import (
	"github.com/gofrs/uuid"
	"time"
)

type Answer struct {
	Id         uuid.UUID `json:"id"`
	QuestionId uuid.UUID `json:"question_id"`
	Body       string    `json:"body"`
	Valid      bool      `json:"valid"`
}

type Course struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	TeacherId  uuid.UUID `json:"teacher_id"`
	SchoolYear int       `json:"school_year"`
}

type Question struct {
	Id      uuid.UUID `json:"id"`
	Body    string    `json:"body"`
	ImgFile string    `json:"img_file"`
	TestId  uuid.UUID `json:"test_id"`
}

type Teacher struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	SecondName string    `json:"second_name"`
	Surname    string    `json:"surname"`
}

type Test struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	CreatedBy string     `json:"created_by"`
	ChangedBy *string    `json:"changed_by"`
	CreatedAt time.Time  `json:"created_at"`
	ChangedAt *time.Time `json:"changed_at"`
	CourseId  uuid.UUID  `json:"course_id"`
}

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Active   bool      `json:"active"`
}
