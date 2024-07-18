package domain

type PerformanceDataRepoStub struct {
	performanceData []PerformanceData
}

func (s PerformanceDataRepoStub) GetData() ([]PerformanceData, error) {
	return s.performanceData, nil
}

func (s PerformanceDataRepoStub) PostData(performanceData []PerformanceData) error {
	// Do nothing for stubb implementation
	return nil
}

func NewBessRepoStub() PerformanceDataRepoStub {
	performanceData := []PerformanceData{
		{Timestamp: "12345", ActivePower: 120.23, ReactivePower: 10.2},
		{Timestamp: "12346", ActivePower: 120.01, ReactivePower: 10.3},
		{Timestamp: "12347", ActivePower: 119.09, ReactivePower: 10.8},
		{Timestamp: "12348", ActivePower: 119.81, ReactivePower: 10.5},
	}
	return PerformanceDataRepoStub{performanceData: performanceData}
}
