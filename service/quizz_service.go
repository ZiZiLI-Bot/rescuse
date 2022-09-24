package service

import (
	"rescues/model"
	"rescues/repository"
)

type QuizzService interface {
	GetById(id int) (*model.Quizz, error)
	GetAll() ([]model.Quizz, error)
	Create(new *model.Quizz) (*model.Quizz, error)
	Update(Quizz model.Quizz) (*model.Quizz, error)
	Delete(id int) error
}

type quizzService struct {
	quizzRepo model.QuizzRepository
}

// --------------------quizz module------------------------

func (s *quizzService) GetById(id int) (*model.Quizz, error) {
	quizz, err := s.quizzRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return quizz, nil
}

func (s *quizzService) GetAll() ([]model.Quizz, error) {
	quizzs, err := s.quizzRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return quizzs, nil
}

func (s *quizzService) Create(new *model.Quizz) (*model.Quizz, error) {
	quizz, err := s.quizzRepo.Create(new)
	if err != nil {
		return nil, err
	}
	return quizz, nil
}

func (s *quizzService) Update(quizz model.Quizz) (*model.Quizz, error) {
	record, err := s.quizzRepo.Update(quizz)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (s *quizzService) Delete(id int) error {
	if err := s.quizzRepo.Delete(id); err != nil {
		return err
	}
	return nil
}


func NewQuizzService() QuizzService {
	quizzRepo := repository.NewquizzRepository()
	return &quizzService{
		quizzRepo: quizzRepo,
	}
}
