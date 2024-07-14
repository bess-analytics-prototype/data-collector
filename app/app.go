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
	bh := BessHandler{service: service.NewDefaultBessService(domain.NewBessRepoStub())}

	router.GET("/bessTestData", bh.GetBessData)
	router.POST("/bessPostTestData", bh.PostBessData)
	http.ListenAndServe(":8080", router)
}
