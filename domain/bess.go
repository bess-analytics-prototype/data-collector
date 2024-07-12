package domain

type Bess struct {
	Id      string `json:"id"`
	Project string `json:"project"`
	Device  string `json:"device"`
	Data    []Data `json:"data"`
}

type Data struct {
	Timestamp string  `json:"time_stamp"`
	Voltage   float64 `json:"voltage"`
	Current   float64 `json:"current"`
}

type BessRepo interface {
	GetData() (Bess, error)
}
