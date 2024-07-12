package service

import "data-collector/domain"

type BessService interface {
	GetAllBessData() (domain.Bess, error)
}

type DefaultBessService struct {
	repo domain.BessRepo
}

func (s DefaultBessService) GetAllBessData() (domain.Bess, error) {
	return s.repo.GetData()
}

func NewDefaultBessService(repo domain.BessRepo) DefaultBessService {
	return DefaultBessService{repo: repo}
}
