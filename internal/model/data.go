// model/data.go

package model

type BESSData struct {
	Timestamp int64   `json:"timestamp"`
	Voltage   float64 `json:"voltage"`
	Current   float64 `json:"current"`
}
