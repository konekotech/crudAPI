package service

import (
	"github.com/konekotech/crudAPI/db"
	"github.com/konekotech/crudAPI/models"
	"log"
)

type PostServiceInterface interface {
	GetModel() (models.Post, error)
	CreateModel(map[string]int) (models.Post, error)
	UpdateModel(map[string]int) (models.Post, error)
	DeleteModel(string) error
}

type PostService struct {
	PostServiceInterface
}

func (s *PostService) GetModel() ([]models.Post, error) {
	var p []models.Post

	if err := db.DB.Find(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (s *PostService) CreateModel(postData map[string]int) (models.Post, error) {

	p := models.Post{}

	p.Weight, _ = postData["weight"]
	p.Height, _ = postData["height"]

	err := db.DB.Create(&p).Error

	return p, err
}

func (s *PostService) UpdateModel(postData map[string]int) (models.Post, error) {
	p := models.Post{}

	err := db.DB.Model(&p).Where("id = ?", postData["id"]).Updates(models.Post{Weight: postData["weight"], Height: postData["height"]}).Error

	if err != nil {
		log.Println(err)
	}

	return p, err
}

func (s *PostService) DeleteModel(id string) error {
	p := models.Post{}

	err := db.DB.Where("id = ?", id).Delete(&p).Error

	if err != nil {
		log.Println(err)
	}

	return err
}
