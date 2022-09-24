package repository

import (
	"time"

	"rescues/model"
	"rescues/infrastructure"
)

type judgeRepository struct{}

func (r *judgeRepository) GetById(id int) (*model.Judge, error) {
	db := infrastructure.GetDB()
	var judge model.Judge
	if err := db.Where("id = ?", id).First(&judge).Error; err != nil {
		return nil, err
	}
	return &judge, nil
}

func (r *judgeRepository) GetAll() ([]model.Judge, error) {
	db := infrastructure.GetDB()
	var judges []model.Judge

	if err := db.Model(&model.Judge{}).Find(&judges).Error; err != nil {
		return nil, err
	}
	return judges, nil
}

func (r *judgeRepository) Create(newJudge *model.Judge) (*model.Judge, error) {
	db := infrastructure.GetDB()
	if err := db.Create(newJudge).Error; err != nil {
		return nil, err
	}
	return newJudge, nil
}

func (r *judgeRepository) Update(newJudge model.Judge) (*model.Judge, error) {
	db := infrastructure.GetDB()

	if err := db.Model(&newJudge).Where("id = ?", newJudge.Id).Updates(newJudge).Error; err != nil {
		return nil, err
	}

	var judge model.Judge
	if err := db.Where("id = ?", newJudge.Id).First(&judge).Error; err != nil {
		return nil, err
	}
	return &judge, nil
}

func (r *judgeRepository) Delete(id int) (error) {
	db := infrastructure.GetDB()

	if err := db.Model(&model.Judge{Id: id}).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func NewJudgeRepository() *judgeRepository {
	return &judgeRepository{}
}