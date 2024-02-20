package server

package server

import (
"github.com/konekotech/crudAPI/controller"
"github.com/konekotech/crudAPI/service
"github.com/gin-gonic/gin"
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
