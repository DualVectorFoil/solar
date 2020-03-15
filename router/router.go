package router

import (
	"github.com/DualVectorFoil/solar/util/jsonUtil"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {

	router := gin.Default()
	router.NoRoute()

	dataRouter := router.Group("/data_service")
	dataRouter.POST("/test", nil)

	router.Run(":13145")
}

func notFound(c *gin.Context) {
	c.String(http.StatusNotFound, jsonUtil.JsonResp(http.StatusNotFound, nil, "404 router not found"))
}
