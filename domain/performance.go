package domain

// type Bess struct {
// 	Id      string `json:"id"`
// 	Project string `json:"project"`
// 	Device  string `json:"device"`
// 	Data    []Data `json:"data"`
// }

// type Data struct {
// 	Timestamp string  `json:"time_stamp"`
// 	Voltage   float64 `json:"voltage"`
// 	Current   float64 `json:"current"`
// }

// type BessRepo interface {
// 	GetData() (Bess, error)
// 	PostData(Bess) error
// }

// Domain object
type PerformanceData struct {
	Timestamp     string  `json:"time_stamp"`
	ActivePower   float64 `json:"active_power"`
	ReactivePower float64 `json:"reactive_power"`
}

// Port
type PerformanceDataRepo interface {
	GetData() ([]PerformanceData, error)
	PostData([]PerformanceData) error
}
