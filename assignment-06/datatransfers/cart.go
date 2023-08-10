package datatransfers

type AddCartItem struct {
	ProductID uint `json:"productId" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required,gt=0"`
}

type Checkout struct {
	TotalPrice float64 `json:"totalPrice"`
}
