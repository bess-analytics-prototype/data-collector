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

	router.GET("/bessTestData", bh.getBessTestData)
	router.POST("/bessPostTestData", bh.postBessData)
	http.ListenAndServe("localhost:8080", router)
}
