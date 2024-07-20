package app

import (
	"data-collector/domain"
	"data-collector/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	// wiring
	var bh BessHandler
	repo, err := domain.NewPerformanceDataRepoDb("PerformanceData")
	if err != nil {
		log.Println("Unable to establish new performance data repo db")
	}

	bh = BessHandler{service: service.NewDefaultPerformanceDataService(repo)}

	router.POST("/postPerformanceTestData", bh.PostPerformanceData)
	http.ListenAndServe(":8080", router)
}
