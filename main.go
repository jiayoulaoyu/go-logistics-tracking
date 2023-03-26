package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-logistics-tracking/consts"
	"go-logistics-tracking/model"
	"go-logistics-tracking/service"
	"io"
	"net/http"
)

func main() {
	r := gin.Default()
	trackRoute(r)
	r.Run("0.0.0.0:7008") // 监听并在 0.0.0.0:8080 上启动服务
}

func trackRoute(r *gin.Engine) {
	r.POST("/logistics/tracking/inner", func(c *gin.Context) {
		arr, _ := io.ReadAll(c.Request.Body)
		var request = new(model.InnerTrackRequest)

		json.Unmarshal(arr, request)
		fmt.Println(string(arr))
		if request.PassWord != consts.PASSWORD {
			c.JSON(http.StatusOK, model.NewErrorResponse(consts.NEGATIVE_7, ""))
			return
		}
		if _, ok := service.API_MAP[request.ApiName]; !ok {
			c.JSON(http.StatusOK, model.NewErrorResponse(consts.ZERO, "接口"+request.ApiName+"不存在"))
			return
		}
		c.JSON(http.StatusOK, model.NewSuccessResponse(service.Execute(request.TrackRequest, request.ApiName)))
	})
}
