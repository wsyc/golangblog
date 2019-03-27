package router

import (
	"gopkg.in/gin-gonic/gin.v1"
	. "newland/controller"
	"net/http"
	. "newland/router/middelware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.POST("/AddPerson", AddPersonApi)
	router.GET("/Login", LoginPersonApi)

	router.GET("/persons", GetPersonsApi)
	router.GET("/news", GetNewsApi)
	router.GET("/types", GetTypesApi)
	//router.GET("/token", SetToken)

	//router.GET("/Token", SetToken)

	//添加群组中间件
	authorized := router.Group("/user", MyMiddelware())
	authorized.GET("/info", func(c *gin.Context) {
		c.String(http.StatusOK, "info")
	})
	//authorized.GET("/login", SetToken)
	//添加群组中间件
	authorizedv2 := router.Group("/v2/user", MyMiddelware())
	authorizedv2.GET("/info", func(c *gin.Context) {
		c.String(http.StatusOK, "info")
	})
	return router
}