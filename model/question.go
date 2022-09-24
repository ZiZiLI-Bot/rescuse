package model

import (
	"time"
)

type Question struct {
	Id            int    `json:"id" gorm:"primaryKey"`
	Question      string `json:"question"`

	QuestionGroup int    `json:"question_group"`

	CreatedAt time.Time  `swaggerignore:"true"`
	UpdatedAt time.Time  `swaggerignore:"true"`
	DeletedAt *time.Time `swaggerignore:"true"`
}

type QuestionRepository interface {
	GetById(id int) (*Question, error)
	GetAll() ([]Question, error)
	Create(new *Question) (*Question, error)
	Update(question Question) (*Question, error)
	Delete(id int) error
	FilterByGroup(groupId int) ([]Question, error)
}
