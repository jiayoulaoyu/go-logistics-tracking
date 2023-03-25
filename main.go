package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-logistics-tracking/model"
	"go-logistics-tracking/service"
	"io"
)

func main() {
	r := gin.Default()
	trackRoute(r)
	fmt.Println()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func trackRoute(r *gin.Engine) {
	r.POST("/client/track", func(c *gin.Context) {
		arr, _ := io.ReadAll(c.Request.Body)
		var request model.TrackRequest
		json.Unmarshal(arr, request)
		res := service.Execute([]model.TrackRequest{request}, "UPS")
		c.JSON(200, res)
	})
}
