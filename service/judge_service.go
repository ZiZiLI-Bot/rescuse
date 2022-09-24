package service

import (
	"rescues/model"
	"rescues/repository"
)

type JudgeService interface {
	GetById(id int) (*model.Judge, error)
	GetAll() ([]model.Judge, error)
	GetAdvices(stress, depess, anxiety int) (*model.AdvicePayload, error)	// nhập điểm của bài đánh giá
	Create(new *model.Judge) (*model.Judge, error)
	Update(judge model.Judge) (*model.Judge, error)
	Delete(id int) error
}

type judgeService struct {
	judgeRepo model.JudgeRepository
}

// --------------------Judge module------------------------

func (s *judgeService) GetById(id int) (*model.Judge, error) {
	judge, err := s.judgeRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return judge, nil
}

func (s *judgeService) GetAll() ([]model.Judge, error) {
	judges, err := s.judgeRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return judges, nil
}

func (s *judgeService) GetAdvices(stress, depess, anxiety int) (*model.AdvicePayload, error) {
	judges, err := s.judgeRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var adviceStress, adviceDepress, adviceAnxiety string
	for _, i := range judges {
		if i.ScoreStressMin <= stress && stress <= i.ScoreStressMax {
			adviceStress = i.AdviceStress
		}
		if i.ScoreDepressMin <= depess && depess <= i.ScoreDepressMax {
			adviceDepress = i.AdviceDepress
		}
		if i.ScoreAnxietyMin <= anxiety && anxiety <= i.ScoreAnxietyMax {
			adviceAnxiety = i.AdviceAnxiety
		}
	}

	return &model.AdvicePayload{
		ScoreStress:  stress,
		ScoreDepress: depess,
		ScoreAnxiety: anxiety,

		AdviceStress:  adviceStress,
		AdviceDepress: adviceDepress,
		AdviceAnxiety: adviceAnxiety,
	}, nil

}

func (s *judgeService) Create(new *model.Judge) (*model.Judge, error) {
	judge, err := s.judgeRepo.Create(new)
	if err != nil {
		return nil, err
	}
	return judge, nil
}

func (s *judgeService) Update(judge model.Judge) (*model.Judge, error) {
	record, err := s.judgeRepo.Update(judge)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (s *judgeService) Delete(id int) error {
	if err := s.judgeRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func NewJudgeService() JudgeService {
	return &judgeService{
		judgeRepo: repository.NewJudgeRepository(),
	}
}
