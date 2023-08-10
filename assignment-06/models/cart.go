package models

import (
	"github.com/ngoduongkha/go-2023/assignment-06/datatransfers"
	"gorm.io/gorm"
)

type cartOrm struct {
	db *gorm.DB
}

type cartItemOrm struct {
	db *gorm.DB
}

type Cart struct {
	ID     uint       `gorm:"primaryKey" json:"id"`
	UserID uint       `gorm:"uniqueIndex" json:"-"`
	Items  []CartItem `json:"items"`
}

type CartItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	CartID    uint    `gorm:"uniqueIndex" json:"-"`
	ProductID uint    `gorm:"uniqueIndex" json:"productId"`
	Quantity  uint    `gorm:"not null" json:"quantity"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
}

type CartOrmer interface {
	GetOneByUserID(userID uint) (cart Cart, err error)
	InsertCartItem(userId uint, dto datatransfers.AddCartItem) (err error)
	DeleteCartItem(userId uint, cartItemId uint) (err error)
	Checkout(userId uint) (response datatransfers.Checkout, err error)
}

type CartItemOrmer interface {
}

func NewCartOrmer(db *gorm.DB) CartOrmer {
	_ = db.AutoMigrate(&Cart{}) // builds table when enabled
	return &cartOrm{db}
}

func NewCartItemOrmer(db *gorm.DB) CartItemOrmer {
	_ = db.AutoMigrate(&CartItem{}) // builds table when enabled
	return &cartItemOrm{db}
}

func (o *cartOrm) GetOneByUserID(userID uint) (cart Cart, err error) {
	result := o.db.Model(&Cart{}).Preload("Items").Where("user_id = ?", userID).First(&cart)
	if result.Error == gorm.ErrRecordNotFound {
		cart = Cart{UserID: userID}
		result = o.db.Model(&Cart{}).Create(&cart)
	}
	return cart, result.Error
}

func (o *cartOrm) InsertCartItem(userId uint, dto datatransfers.AddCartItem) (err error) {
	cart, err := o.GetOneByUserID(userId)
	if err == gorm.ErrRecordNotFound {
		cart = Cart{UserID: userId}
		o.db.Model(&Cart{}).Create(&cart)
	}
	existingCartItem := CartItem{}
	result := o.db.Model(&CartItem{}).Where("cart_id = ? AND product_id = ?", cart.ID, dto.ProductID).First(&existingCartItem)
	if result.Error == gorm.ErrRecordNotFound {
		cartItem := CartItem{CartID: cart.ID, ProductID: dto.ProductID, Quantity: dto.Quantity}
		result = o.db.Model(&CartItem{}).Create(&cartItem)
	} else {
		result = o.db.Model(&existingCartItem).Update("quantity", dto.Quantity)
	}
	return result.Error
}

func (o *cartOrm) DeleteCartItem(userId uint, cartItemId uint) (err error) {
	cart, err := o.GetOneByUserID(userId)
	if err == gorm.ErrRecordNotFound {
		return
	}
	cartItem := CartItem{}
	result := o.db.Model(&CartItem{}).Where("cart_id = ? AND id = ?", cart.ID, cartItemId).First(&cartItem)
	if result.Error == gorm.ErrRecordNotFound {
		return
	}
	result = o.db.Model(&cartItem).Delete(&cartItem)
	return result.Error
}

func (o *cartOrm) Checkout(userId uint) (response datatransfers.Checkout, err error) {
	res := datatransfers.Checkout{
		TotalPrice: 0,
	}
	// load cart then items and calculate total price
	cart, err := o.GetOneByUserID(userId)
	if err == gorm.ErrRecordNotFound {
		return
	}
	cartItems := []CartItem{}
	result := o.db.Model(&CartItem{}).Preload("Product").Where("cart_id = ?", cart.ID).Find(&cartItems)
	if result.Error == gorm.ErrRecordNotFound {
		return
	}
	for _, cartItem := range cartItems {
		res.TotalPrice += float64(cartItem.Quantity) * cartItem.Product.Price
	}
	// delete cart and cart items
	o.db.Model(&cart).Delete(&cart)
	o.db.Model(&cartItems).Delete(&cartItems)
	return res, result.Error
}
