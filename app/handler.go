package app

import (
	"data-collector/domain"
	"data-collector/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type BessHandler struct {
	service service.BessService
}

func (bh *BessHandler) GetBessData(c *gin.Context) {
	bess, err := bh.service.GetAllBessData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch BESS data"})
		return
	}
	c.JSON(http.StatusOK, bess)
}

func (bh *BessHandler) PostBessData(c *gin.Context) {
	var bess domain.Bess
	if err := c.ShouldBindJSON(&bess); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := bh.service.PostAllBessData(bess)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data saved successfully"})
}
