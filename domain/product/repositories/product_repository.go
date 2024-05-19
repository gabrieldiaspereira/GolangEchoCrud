package repositories

import (
	"github.com/gabrieldiaspereira/echoGoApi/domain/product/models"

	"gorm.io/gorm"
)

type dbProduct struct {
	Conn *gorm.DB
}

func (db *dbProduct) Create(product models.Product) (models.Product, error) {
	result := db.Conn.Create(&product)
	return product, result.Error
}

func (db *dbProduct) Delete(idProduct int) error {
	return db.Conn.Delete(&models.Product{Id: idProduct}).Error
}

func (db *dbProduct) GetAll() ([]models.Product, error) {
	var data []models.Product
	result := db.Conn.Find(&data)
	return data, result.Error
}

func (db *dbProduct) GetById(idProduct int) (models.Product, error) {
	var data models.Product
	result := db.Conn.Where("id", idProduct).First(&data)
	return data, result.Error
}

func (db *dbProduct) Update(idProduct int, product models.Product) (models.Product, error) {
	var data models.Product
	result := db.Conn.Where("id", idProduct).Updates(product).First(&data)
	return data, result.Error
}

type ProductRepository interface {
	Create(product models.Product) (models.Product, error)
	Update(idProduct int, product models.Product) (models.Product, error)
	Delete(idProduct int) error
	GetById(idProduct int) (models.Product, error)
	GetAll() ([]models.Product, error)
}

func NewProductRepository(Conn *gorm.DB) ProductRepository {
	return &dbProduct{Conn: Conn}
}
