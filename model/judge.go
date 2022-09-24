package model

import (
	"time"
)

type Judge struct {
	Id int `json:"id" gorm:"primaryKey"`

	// Phạm vi điểm số
	ScoreStressMax  int `json:"score_stress_max"`
	ScoreStressMin  int `json:"score_stress_min"`
	ScoreDepressMax int `json:"score_depess_max"`
	ScoreDepressMin int `json:"score_depess_min"`
	ScoreAnxietyMax int `json:"score_anxiety_max"`
	ScoreAnxietyMin int `json:"score_anxiety_min"`

	// Lời khuyên tương ứng
	AdviceStress  string `json:"advice_stress"`
	AdviceDepress string `json:"advice_depess"`
	AdviceAnxiety string `json:"advice_anxiety"`

	CreatedAt time.Time  `swaggerignore:"true"`
	UpdatedAt time.Time  `swaggerignore:"true"`
	DeletedAt *time.Time `swaggerignore:"true"`
}

type AdvicePayload struct {
	AdviceStress  string `json:"advice_stress"`
	AdviceDepress string `json:"advice_depess"`
	AdviceAnxiety string `json:"advice_anxiety"`

	ScoreStress  int `json:"score_stress"`
	ScoreDepress int `json:"score_depess"`
	ScoreAnxiety int `json:"score_anxiety"`
}

type JudgeRepository interface {
	GetById(id int) (*Judge, error)
	// GetByScore(stress, depess, anxiety int) (*Judge, error)
	GetAll() ([]Judge, error)
	Create(new *Judge) (*Judge, error)
	Update(judge Judge) (*Judge, error)
	Delete(id int) error
}
