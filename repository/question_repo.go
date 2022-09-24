package repository

import (
	"time"

	"rescues/infrastructure"
	"rescues/model"
)

type questionRepository struct{}

func (r *questionRepository) GetById(id int) (*model.Question, error) {
	db := infrastructure.GetDB()
	var question model.Question
	if err := db.Where("id = ?", id).First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *questionRepository) GetAll() ([]model.Question, error) {
	db := infrastructure.GetDB()
	var questions []model.Question

	if err := db.Model(&model.Question{}).Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}


func (r *questionRepository) Create(newQuestion *model.Question) (*model.Question, error) {
	db := infrastructure.GetDB()
	if err := db.Create(newQuestion).Error; err != nil {
		return nil, err
	}
	return newQuestion, nil
}

func (r *questionRepository) Update(newQuestion model.Question) (*model.Question, error) {
	db := infrastructure.GetDB()

	if err := db.Model(&newQuestion).Where("id = ?", newQuestion.Id).Updates(newQuestion).Error; err != nil {
		return nil, err
	}

	var Question model.Question
	if err := db.Where("id = ?", newQuestion.Id).First(&Question).Error; err != nil {
		return nil, err
	}
	return &Question, nil
}

func (r *questionRepository) Delete(id int) (error) {
	db := infrastructure.GetDB()

	if err := db.Model(&model.Question{Id: id}).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func (r *questionRepository) FilterByGroup(groupId int) ([]model.Question, error) {
	db := infrastructure.GetDB()
	var questions []model.Question

	if err := db.Where("question_group = ?", groupId).Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func NewQuestionRepository() model.QuestionRepository {
	return &questionRepository{}
}