package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/basicStartupTestData", getBasicStartupTestData).Methods(http.MethodGet)
	http.ListenAndServe("localhost:8080", router)
}
