package controllers

import (
	"net/http"
	"strconv"

	"github.com/gabrieldiaspereira/echoGoApi/domain/product/models"
	"github.com/gabrieldiaspereira/echoGoApi/domain/product/services"

	vl "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductController struct {
	productService services.ProductService
	validate       vl.Validate
}

func (controller ProductController) Create(c echo.Context) error {
	type payload struct {
		Name        string  `json:"name" validate:"required"`
		Description string  `json:"description" validate:"required"`
		Price       float64 `json:"price" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return err
	}

	result := controller.productService.Create(models.Product{Name: payloadValidator.Name, Description: payloadValidator.Description, Price: payloadValidator.Price})

	return c.JSON(http.StatusOK, result)
}

func (controller ProductController) GetAll(c echo.Context) error {
	result := controller.productService.GetAll()
	return c.JSON(http.StatusOK, result)
}

func (controller ProductController) GetById(c echo.Context) error {
	idProduct, _ := strconv.Atoi(c.Param("id_product"))
	result := controller.productService.GetById(idProduct)
	if result.Id <= 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, result)
}

func (controller ProductController) Update(c echo.Context) error {
	type payload struct {
		Name        string  `json:"name" validate:"required"`
		Description string  `json:"description" validate:"required"`
		Price       float64 `json:"price" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	idProduct, _ := strconv.Atoi(c.Param("id_product"))
	result := controller.productService.Update(idProduct, models.Product{Name: payloadValidator.Name, Description: payloadValidator.Description, Price: payloadValidator.Price})
	if result.Id <= 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, result)
}

func (controller ProductController) Delete(c echo.Context) error {
	idProduct, _ := strconv.Atoi(c.Param("id_product"))
	result := controller.productService.Delete(idProduct)
	return c.JSON(http.StatusOK, result)
}

func NewProductController(db *gorm.DB) ProductController {
	service := services.NewProductService(db)
	controller := ProductController{
		productService: service,
		validate:       *vl.New(),
	}

	return controller
}
