package app

import (
	"data-collector/domain"
	"data-collector/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	//wiring
	var bh BessHandler
	repo, err := domain.NewBessRepoDb("someTable")
	if err != nil {
		bh = BessHandler{service: service.NewDefaultBessService(domain.NewBessRepoStub())}
	} else {
		bh = BessHandler{service: service.NewDefaultBessService(repo)}
	}

	router.GET("/bessTestData", bh.GetBessData)
	router.POST("/bessPostTestData", bh.PostBessData)
	http.ListenAndServe(":8080", router)
}
