package service

import "data-collector/domain"

type PerformanceService interface {
	PostAllPerformanceData([]domain.PerformanceData) error
}

type DefaultPerformanceDataService struct {
	repo domain.PerformanceDataRepo
}

func (s DefaultPerformanceDataService) PostAllPerformanceData(performanceData []domain.PerformanceData) error {
	return s.repo.PostData(performanceData)
}

func NewDefaultPerformanceDataService(repo domain.PerformanceDataRepo) DefaultPerformanceDataService {
	return DefaultPerformanceDataService{repo: repo}
}
