package app

import (
	"data-collector/domain"
	"data-collector/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	bh := BessHandler{service: service.NewDefaultBessService(domain.NewBessRepoStub())}

	router.HandleFunc("/bessTestData", bh.getBessTestData).Methods(http.MethodGet)
	router.HandleFunc("/bessPostTestData", bh.postBessData).Methods(http.MethodPost)
	http.ListenAndServe("localhost:8080", router)
}
