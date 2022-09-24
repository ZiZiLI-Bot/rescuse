package repository

import (
	"time"

	"rescues/infrastructure"
	"rescues/model"
)

type quizzRepository struct{}

func (r *quizzRepository) GetById(id int) (*model.Quizz, error) {
	db := infrastructure.GetDB()
	var quizz model.Quizz
	if err := db.Where("id = ?", id).First(&quizz).Error; err != nil {
		return nil, err
	}
	return &quizz, nil
}

func (r *quizzRepository) GetAll() ([]model.Quizz, error) {
	db := infrastructure.GetDB()
	var quizzs []model.Quizz

	if err := db.Model(&model.Quizz{}).Find(&quizzs).Order("id ASC").Error; err != nil {
		return nil, err
	}
	return quizzs, nil
}

func (r *quizzRepository) Create(newquizz *model.Quizz) (*model.Quizz, error) {
	db := infrastructure.GetDB()
	if err := db.Create(newquizz).Error; err != nil {
		return nil, err
	}
	return newquizz, nil
}

func (r *quizzRepository) Update(newquizz model.Quizz) (*model.Quizz, error) {
	db := infrastructure.GetDB()

	if err := db.Model(&newquizz).Where("id = ?", newquizz.Id).Updates(newquizz).Error; err != nil {
		return nil, err
	}

	var quizz model.Quizz
	if err := db.Where("id = ?", newquizz.Id).First(&quizz).Error; err != nil {
		return nil, err
	}
	return &quizz, nil
}

func (r *quizzRepository) Delete(id int) error {
	db := infrastructure.GetDB()

	if err := db.Model(&model.Quizz{Id: id}).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func NewquizzRepository() model.QuizzRepository {
	return &quizzRepository{}
}
