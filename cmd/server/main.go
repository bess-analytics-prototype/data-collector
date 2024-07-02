// main.go

package main

import (
	"data-collector/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/data", api.PostData)
	r.Run(":8080") // listen on port 8080
}
