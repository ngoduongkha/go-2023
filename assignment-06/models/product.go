package models

import (
	"time"

	"gorm.io/gorm"
)

type productOrm struct {
	db *gorm.DB
}

type Product struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"uniqueIndex" json:"name"`
	Price     float64   `json:"price"`
	Quantity  uint      `json:"quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

type ProductOrmer interface {
	GetAll() (products []Product, err error)
	GetOneByID(id uint) (product Product, err error)
	InsertProduct(product Product) (err error)
	UpdateProduct(product Product) (err error)
	DeleteByID(id uint) (err error)
}

func NewProductOrmer(db *gorm.DB) ProductOrmer {
	_ = db.AutoMigrate(&Product{}) // builds table when enabled
	return &productOrm{db}
}

func (o *productOrm) GetAll() (products []Product, err error) {
	result := o.db.Model(&Product{}).Find(&products)
	return products, result.Error
}

func (o *productOrm) GetOneByID(id uint) (product Product, err error) {
	result := o.db.Model(&Product{}).Where("id = ?", id).First(&product)
	return product, result.Error
}

func (o *productOrm) InsertProduct(product Product) (err error) {
	result := o.db.Model(&Product{}).Create(&product)
	return result.Error
}

func (o *productOrm) UpdateProduct(product Product) (err error) {
	result := o.db.Model(&Product{}).Model(&product).Updates(&product)
	return result.Error
}

func (o *productOrm) DeleteByID(id uint) (err error) {
	result := o.db.Model(&Product{}).Where("id = ?", id).Delete(&Product{})
	return result.Error
}
