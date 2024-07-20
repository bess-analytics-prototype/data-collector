package app

import (
	"data-collector/domain"
	"data-collector/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type BessHandler struct {
	service service.PerformanceService
}

func (bh *BessHandler) PostPerformanceData(c *gin.Context) {
	var performanceData []domain.PerformanceData
	if err := c.ShouldBindJSON(&performanceData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := bh.service.PostAllPerformanceData(performanceData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data saved successfully"})
}
