package domain

// stub adapter
type BessRepoStub struct {
	bess Bess
}

func (s BessRepoStub) GetData() (Bess, error) {
	return s.bess, nil
}

func (s BessRepoStub) PostData(bess Bess) error {
	// Do nothing until db is built and configured
	return nil
}

func NewBessRepoStub() BessRepoStub {
	bess := Bess{
		Id:      "1234",
		Project: "SomeProject",
		Device:  "BESS",
		Data: []Data{
			{Timestamp: "12345", Voltage: 120.23, Current: 10.2},
			{Timestamp: "12346", Voltage: 120.01, Current: 10.3},
			{Timestamp: "12347", Voltage: 119.09, Current: 10.8},
			{Timestamp: "12348", Voltage: 119.81, Current: 10.5},
		},
	}
	return BessRepoStub{bess: bess}
}
