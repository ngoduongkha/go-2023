package datatransfers

type ProductCreate struct {
	Name     string  `json:"name" binding:"required,max=255,min=3"`
	Price    float64 `json:"price" binding:"required,gt=0"`
	Quantity uint    `json:"quantity" binding:"required,gt=0"`
}

type ProductUpdate struct {
	Name     string  `json:"name" binding:"required,max=255,min=3"`
	Price    float64 `json:"price" binding:"required,gt=0"`
	Quantity uint    `json:"quantity" binding:"required,gt=0"`
}
