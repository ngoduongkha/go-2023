package handlers

import (
	"errors"

	"github.com/ngoduongkha/go-2023/assignment-06/datatransfers"
	"github.com/ngoduongkha/go-2023/assignment-06/models"
)

func (m *module) GetProducts() (products []models.Product, err error) {
	if products, err = m.db.productOrmer.GetAll(); err != nil {
		return products, errors.New("cannot get all products")
	}
	return
}

func (m *module) GetProduct(id uint) (product models.Product, err error) {
	if product, err = m.db.productOrmer.GetOneByID(id); err != nil {
		return product, errors.New("cannot retrieve product")
	}
	return
}

func (m *module) UpdateProduct(id uint, product datatransfers.ProductUpdate) (err error) {
	if err = m.db.productOrmer.UpdateProduct(models.Product{
		ID:       id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}); err != nil {
		return errors.New("cannot update product")
	}
	return
}

func (m *module) CreateProduct(product datatransfers.ProductCreate) (err error) {
	if err := m.db.productOrmer.InsertProduct(models.Product{
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}); err != nil {
		return errors.New("cannot create product")
	}
	return
}
