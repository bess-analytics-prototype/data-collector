package domain

// Domain object
type PerformanceData struct {
	Timestamp     string  `json:"time_stamp"`
	ActivePower   float64 `json:"active_power"`
	ReactivePower float64 `json:"reactive_power"`
}

// Port
type PerformanceDataRepo interface {
	PostData([]PerformanceData) error
}
