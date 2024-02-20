package server

import (
	"github.com/gin-gonic/gin"
	"github.com/konekotech/crudAPI/controller"
	"github.com/konekotech/crudAPI/service"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	s := service.PostService{}
	c := controller.NewPostController(s)
	r.GET("", c.Index)
	r.POST("/post", c.Create)
	r.PUT("/:id", c.Update)
	r.DELETE("/:id", c.Delete)

	return r
}
