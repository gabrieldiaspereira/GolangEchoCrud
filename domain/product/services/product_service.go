package services

import (
	"fmt"

	"github.com/gabrieldiaspereira/echoGoApi/domain/product/models"
	"github.com/gabrieldiaspereira/echoGoApi/domain/product/repositories"
	"github.com/gabrieldiaspereira/echoGoApi/helpers"

	"gorm.io/gorm"
)

type productService struct {
	productRepo repositories.ProductRepository
}

func (service *productService) Create(product models.Product) models.Product {
	var response helpers.Response
	data, err := service.productRepo.Create(product)
	if err != nil {
		response.Status = 500
		response.Messages = "Failed to get products"
	}
	return data
}

func (service *productService) GetAll() []models.Product {
	var response helpers.Response
	data, err := service.productRepo.GetAll()
	if err != nil {
		response.Status = 500
		response.Messages = "Failed to get products"
	}
	return data
}

func (service *productService) GetById(idProduct int) models.Product {
	var response helpers.Response
	data, err := service.productRepo.GetById(idProduct)
	if err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to get product: ", idProduct)
	}

	return data
}

func (service *productService) Update(idProduct int, product models.Product) models.Product {
	var response helpers.Response
	data, err := service.productRepo.Update(idProduct, product)
	if err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to get product: ", idProduct)
	}
	return data
}

func (service *productService) Delete(idProduct int) helpers.Response {
	var response helpers.Response
	if err := service.productRepo.Delete(idProduct); err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to delete product: ", idProduct)
	} else {
		response.Status = 200
		response.Messages = "Success to delete product"
	}
	return response
}

type ProductService interface {
	Create(product models.Product) models.Product
	GetAll() []models.Product
	GetById(idProduct int) models.Product
	Update(idProduct int, product models.Product) models.Product
	Delete(idProduct int) helpers.Response
}

func NewProductService(db *gorm.DB) ProductService {
	return &productService{productRepo: repositories.NewProductRepository(db)}
}
