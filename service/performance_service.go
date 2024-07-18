package service

import "data-collector/domain"

type PerformanceService interface {
	GetAllPerformanceData() ([]domain.PerformanceData, error)
	PostAllPerformanceData([]domain.PerformanceData) error
}

type DefaultPerformanceDataService struct {
	repo domain.PerformanceDataRepo
}

func (s DefaultPerformanceDataService) GetAllPerformanceData() ([]domain.PerformanceData, error) {
	return s.repo.GetData()
}

func (s DefaultPerformanceDataService) PostAllPerformanceData(performanceData []domain.PerformanceData) error {
	return s.repo.PostData(performanceData)
}

func NewDefaultPerformanceDataService(repo domain.PerformanceDataRepo) DefaultPerformanceDataService {
	return DefaultPerformanceDataService{repo: repo}
}
