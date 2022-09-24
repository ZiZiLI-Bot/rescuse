package service

import (
	"rescues/model"
	"rescues/repository"
)

type QuestionService interface {
	GetById(id int) (*model.Question, error)
	GetAll() ([]model.Question, error)
	Create(new *model.Question) (*model.Question, error)
	Update(question model.Question) (*model.Question, error)
	Delete(id int) error
	FilterByGroup(groupId int) ([]model.Question, error)
}

type questionService struct {
	questionRepo model.QuestionRepository
}

// --------------------question module------------------------

func (s *questionService) GetById(id int) (*model.Question, error) {
	question, err := s.questionRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (s *questionService) GetAll() ([]model.Question, error) {
	questions, err := s.questionRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (s *questionService) Create(new *model.Question) (*model.Question, error) {
	question, err := s.questionRepo.Create(new)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (s *questionService) Update(question model.Question) (*model.Question, error) {
	record, err := s.questionRepo.Update(question)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (s *questionService) Delete(id int) error {
	if err := s.questionRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (s *questionService) FilterByGroup(groupId int) ([]model.Question, error) {
	questions, err := s.questionRepo.FilterByGroup(groupId)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func NewquestionService() QuestionService {
	questionRepo := repository.NewQuestionRepository()
	return &questionService{
		questionRepo: questionRepo,
	}
}
