package handlers

import (
	"errors"
	"github.com/ngoduongkha/go-2023/assignment-06/datatransfers"
	"github.com/ngoduongkha/go-2023/assignment-06/models"
)

func (m *module) GetCart(userId uint) (cart models.Cart, err error) {
	if cart, err = m.db.cartOrmer.GetOneByUserID(userId); err != nil {
		return cart, errors.New("cannot retrieve cart")
	}
	return
}

func (m *module) AddToCart(userId uint, dto datatransfers.AddCartItem) (err error) {
	if err = m.db.cartOrmer.InsertCartItem(userId, dto); err != nil {
		return errors.New("cannot add to cart")
	}
	return
}

func (m *module) RemoveFromCart(userId uint, cartItemId uint) (err error) {
	if err = m.db.cartOrmer.DeleteCartItem(userId, cartItemId); err != nil {
		return errors.New("cannot delete cart item")
	}
	return
}

func (m *module) Checkout(userId uint) (datatransfers.Checkout, error) {
	res, err := m.db.cartOrmer.Checkout(userId)
	if err != nil {
		return datatransfers.Checkout{}, errors.New("cannot checkout")
	}
	return res, nil
}
