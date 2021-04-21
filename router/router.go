package router

import (
	"gin-app/config"
	v1 "gin-app/controller/api/v1"
	"gin-app/middleware/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	// debug or release or test
	if config.Bases.App.Debug {
		router.Use(gin.Logger())
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	apiv1 := router.Group("/api/v1")

	// 跳出三界之外
	apiv1.GET("/get_token", v1.Auth.GetAuth)

	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/demo", v1.Demo.Index)
		apiv1.GET("/demo/verify", v1.Demo.Verify)
	}

	router.GET("/v", func(c *gin.Context) {
		e := config.Get("RonyLee")

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  e,
		})
	})

	return router
}
