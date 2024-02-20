package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/konekotech/crudAPI/service"
	"strconv"
)

type PostController struct {
	service.PostService
}

func NewPostController(s service.PostService) *PostController {
	return &PostController{
		PostService: s,
	}
}

func (controller *PostController) Index(c *gin.Context) {
	p, err := controller.GetModel()

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (controller *PostController) Create(c *gin.Context) {
	weight, _ := strconv.Atoi(c.PostForm("Weight"))
	height, _ := strconv.Atoi(c.PostForm("Height"))

	postData := map[string]int{
		"weight": weight,
		"height": height,
	}

	p, err := controller.CreateModel(postData)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (controller *PostController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	weight, _ := strconv.Atoi(c.PostForm("Weight"))
	height, _ := strconv.Atoi(c.PostForm("Height"))

	postData := map[string]int{
		"id":     id,
		"weight": weight,
		"height": height,
	}

	p, err := controller.UpdateModel(postData)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (controller *PostController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	err := controller.DeleteModel(id)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, "OK")
	}
}
