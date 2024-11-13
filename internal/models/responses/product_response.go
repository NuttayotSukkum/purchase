package responses

type CreateProductResp struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    ProductResp `json:"data"`
}

type CreateProductIdResp struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    ProductGetIdResp `json:"data"`
}

type ProductListResp struct {
	Code    string             `json:"code"`
	Message string             `json:"message"`
	Data    []ProductGetIdResp `json:"data"`
}

type ProductResp struct {
	Id     string
	Name   string
	Amount int
	Price  float64
}

type ProductGetIdResp struct {
	Id     string
	Name   string
	Amount int
	Price  float64
	Status string
}

func (resp *ProductResp) BuildProductResp() CreateProductResp {
	return CreateProductResp{
		Code:    "1000",
		Message: "Success",
		Data: ProductResp{
			Id:     resp.Id,
			Name:   resp.Name,
			Amount: resp.Amount,
			Price:  resp.Price,
		},
	}
}

func (resp *ProductGetIdResp) BuildProductIdResp() CreateProductIdResp {
	return CreateProductIdResp{
		Code:    "1000",
		Message: "Success",
		Data: ProductGetIdResp{
			Id:     resp.Id,
			Name:   resp.Name,
			Amount: resp.Amount,
			Price:  resp.Price,
			Status: resp.Status,
		},
	}
}

//func (resp *ProductGetIdResp) BuildProductListResp() ProductListResp {
//	return ProductListResp{
//		Code:    "1000",
//		Message: "Success",
//		Data: []ProductGetIdResp{
//			{
//				Id:     resp.Id,
//				Name:   resp.Name,
//				Amount: resp.Amount,
//				Price:  resp.Price,
//				Status: resp.Status,
//			},
//		},
//	}
//}
