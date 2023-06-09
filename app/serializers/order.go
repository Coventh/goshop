package serializers

type Order struct {
	ID         string       `json:"id"`
	Code       string       `json:"code"`
	Lines      []*OrderLine `json:"lines"`
	TotalPrice uint         `json:"total_price"`
	Status     string       `json:"status"`
}

type OrderLine struct {
	Product  Product `json:"product,omitempty"`
	Quantity uint    `json:"quantity"`
	Price    uint    `json:"price"`
}

type PlaceOrderReq struct {
	Lines []PlaceOrderLineReq `json:"lines,omitempty" validate:"required,gt=0,lte=5"`
}

type PlaceOrderLineReq struct {
	ProductID string `json:"product_id,omitempty" validate:"required"`
	Quantity  uint   `json:"quantity,omitempty" validate:"required"`
}

type OrderQueryParam struct {
	Code   string `json:"code,omitempty" form:"code"`
	Status string `json:"status,omitempty" form:"active"`
}
