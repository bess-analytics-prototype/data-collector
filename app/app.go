package app

import (
	"data-collector/domain"
	"data-collector/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	// wiring
	var bh BessHandler
	repo, err := domain.NewPerformanceDataRepoDb("PerformanceData")
	if err != nil {
		bh = BessHandler{service: service.NewDefaultPerformanceDataService(domain.NewBessRepoStub())}
	} else {
		bh = BessHandler{service: service.NewDefaultPerformanceDataService(repo)}
	}

	router.GET("/getPerformanceTestData", bh.GetPerformanceData)
	router.POST("/postPerformanceTestData", bh.PostPerformanceData)
	http.ListenAndServe(":8080", router)
}
