package app

import (
	"data-collector/domain"
	"data-collector/service"
	"encoding/json"
	"net/http"
)

type BessHandler struct {
	service service.BessService
}

func (bh *BessHandler) getBessTestData(w http.ResponseWriter, r *http.Request) {

	bess, _ := bh.service.GetAllBessData()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bess)
}

func (bh *BessHandler) postBessData(w http.ResponseWriter, r *http.Request) {
	var bess domain.Bess
	if err := json.NewDecoder(r.Body).Decode(&bess); err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		err := bh.service.PostAllBessData(bess)
		if err != nil {
			panic(err) //Refine error handling
		}
		json.NewEncoder(w).Encode("Data saved successfully")
	}
}
