package responses

type PaymentIdResponse struct {
	PaymentId string `json:"paymentId"`
}

type CreatePaymentResponse struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Data    PaymentIdResponse `json:"data"`
}

//func (req *CreateProductResp) BuildCreatePaymentResponse() CreatePaymentResponse {
//	return CreatePaymentResponse{
//		Code:    "1000",
//		Message: "success",
//		Data: PaymentIdResponse{
//			PaymentId: req.Data.Id,
//		},
//	}
//}
