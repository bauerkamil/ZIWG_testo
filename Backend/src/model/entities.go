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

type Teacher struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	SecondName string    `json:"second_name"`
	Surname    string    `json:"surname"`
}

type Course struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	UsosId     string    `json:"usos_id"`
	CourseType string    `json:"course_type"`
	TeacherId  uuid.UUID `json:"teacher_id"`
	Teacher    *Teacher  `json:"teacher"  gorm:"foreignKey:TeacherId"`
}

type Question struct {
	Id      uuid.UUID `json:"id"`
	Body    string    `json:"body"`
	ImgFile string    `json:"img_file"`
	Answers []Answer  `json:"answers"`
	TestId  uuid.UUID `json:"test_id"`
}

type Test struct {
	Id         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	CreatedBy  string     `json:"created_by"`
	ChangedBy  *string    `json:"changed_by"`
	CreatedAt  time.Time  `json:"created_at"`
	Course     Course     `json:"course" gorm:"foreignKey:CourseId"`
	CourseId   uuid.UUID  `json:"course_id"`
	ChangedAt  *time.Time `json:"changed_at"`
	Questions  []Question `json:"questions"`
	SchoolYear string     `json:"school_year"`
}
