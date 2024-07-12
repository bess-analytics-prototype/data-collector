package app

import (
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
