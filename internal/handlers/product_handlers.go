package handlers

import (
	"errors"
	"github.com/NuttayotSukkum/purchase/internal/constants"
	dto2 "github.com/NuttayotSukkum/purchase/internal/models/dto"
	"github.com/NuttayotSukkum/purchase/internal/models/entities"
	"github.com/NuttayotSukkum/purchase/internal/models/requests"
	"github.com/NuttayotSukkum/purchase/internal/models/responses"
	"github.com/NuttayotSukkum/purchase/internal/services"
	"github.com/NuttayotSukkum/purchase/internal/utils"
	"github.com/labstack/echo/v4"
	logger "github.com/labstack/gommon/log"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type productHandler struct {
	productManagementService services.ProductService
	paymentService           services.PaymentService
}

func NewProductHandler(productManagementService services.ProductService, paymentService services.PaymentService) *productHandler {
	return &productHandler{
		productManagementService: productManagementService,
		paymentService:           paymentService,
	}
}

func (h *productHandler) CreateProductHandler(c echo.Context) error {
	var product requests.ProductRequest
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, dto2.GenericError)

	}
	logger.Infof("Bound product data: %+v", product)
	if utils.CheckProduct(&product, nil, constants.CHECK_PRODUCT) {
		return c.JSON(http.StatusOK, dto2.GenericError)
	}
	logger.Errorf("Checking ... %s", constants.CHECK_PRODUCT)
	existProduct, err := h.productManagementService.FindProductByName(product)
	if utils.CheckProduct(&product, existProduct, constants.CHECK_PRODUCT_EXIST) {
		logger.Warnf("Product already exists: Name =%s", product.Name)
		return c.JSON(http.StatusBadRequest, dto2.GenericError)

	}
	logger.Errorf("Checking ... ")
	createProductRes, err := h.productManagementService.CreateProduct(product)
	if err != nil {
		logger.Errorf("Unable to fething Data from Database:%v", err)
		return c.JSON(http.StatusInternalServerError, dto2.GenericError)
	}
	productResp := responses.ProductResp{
		Id:     createProductRes.ID,
		Name:   createProductRes.Name,
		Amount: createProductRes.Amount,
		Price:  createProductRes.Price}
	return c.JSON(http.StatusOK, productResp.BuildProductResp())

}

func (h *productHandler) EditProductHandler(c echo.Context) error {
	var req requests.ProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto2.GenericError)
	}
	if utils.CheckProduct(&req, nil, constants.CHECK_PRODUCT) {
		return c.JSON(http.StatusBadRequest, dto2.GenericError)
	}
	productUpdate, err := h.productManagementService.EditProduct(&req.Name, req.Amount, req.Price)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto2.GenericError)
	}
	logger.Infof("Product is update:%+v", productUpdate)
	return c.JSON(http.StatusOK, dto2.GenericSuccess)
}

func (h *productHandler) GetProductHandler(c echo.Context) error {
	var req requests.IdRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto2.GenericError)
	}
	product := h.productManagementService.FindProductByProductId(req.Id)
	if product.Name == "" && product.ID == "" {
		return c.JSON(http.StatusBadRequest, dto2.GenericError)
	}
	if product.Amount <= 0 {
		response := responses.ProductGetIdResp{
			Id:     product.ID,
			Name:   product.Name,
			Amount: product.Amount,
			Price:  product.Price,
			Status: "sold_out",
		}
		return c.JSON(http.StatusOK, response.BuildProductIdResp())
	}
	response := responses.ProductGetIdResp{
		Id:     product.ID,
		Name:   product.Name,
		Amount: product.Amount,
		Price:  product.Price,
		Status: "available",
	}
	return c.JSON(http.StatusOK, response.BuildProductIdResp())
}

func (h *productHandler) PartialSearchProduct(c echo.Context) error {
	var reqProduct requests.ProductNameRequest
	if err := c.Bind(&reqProduct); err != nil {
		return c.JSON(http.StatusBadRequest, dto2.GenericError)
	}
	listOfProduct, err := h.productManagementService.SearchProductByName(reqProduct.ProductName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto2.GenericError)
	}
	productResp := responses.ProductListResp{
		Code:    "1000",
		Message: "Success",
		Data:    *listOfProduct,
	}
	return c.JSON(http.StatusOK, productResp)
}

func (h *productHandler) CreatePayment(c echo.Context) error {
	var reqPayment requests.PaymentRequest
	if err := c.Bind(&reqPayment); err != nil {
		return c.JSON(http.StatusBadRequest, dto2.GenericError)
	}
	if reqPayment.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, dto2.GenericError)
	}
	product := h.productManagementService.FindProductByProductId(reqPayment.ProductId)
	logger.Infof("product: %+v", product)
	if 0 >= product.Amount {
		return c.JSON(http.StatusOK, dto2.ProductNotEnoughError)
	}
	paymentResponse, err := h.paymentService.CreatePayment(reqPayment, product.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto2.GenericError)
	}
	response := responses.PaymentIdResponse{PaymentId: paymentResponse.ID}
	return c.JSON(http.StatusOK, responses.CreatePaymentResponse{
		Code:    "1000",
		Message: "success",
		Data:    response,
	})
}

func (h *productHandler) ConfirmPurchaseHandler(c echo.Context) error {
	var reqPaymentId requests.PaymentIdRequest
	if err := c.Bind(&reqPaymentId); err != nil {
		return c.JSON(http.StatusBadRequest, dto2.GenericError)
	}

	payment, product, err := h.fetchPaymentAndProduct(reqPaymentId)
	if err != nil {
		if errors.Is(err, &dto2.PaymentNotFound) || errors.Is(err, &dto2.ProductNotFound) {
			return c.JSON(http.StatusNotFound, dto2.GenericError)
		}
		return c.JSON(http.StatusInternalServerError, dto2.GenericError)
	}

	if payment.PaymentStatus == constants.SUCCESS_STATUS {
		return c.JSON(http.StatusOK, dto2.GenericSuccess)
	}

	if payment.PaymentStatus != constants.PENDING_STATUS {
		h.paymentService.UpdatePaymentStatus(reqPaymentId.PaymentId, constants.FAILED_STATUS)
		return c.JSON(http.StatusBadRequest, dto2.ProductInvalid)
	}

	if product.Amount <= 0 || payment.Amount > product.Amount {
		h.paymentService.UpdatePaymentStatus(reqPaymentId.PaymentId, constants.FAILED_STATUS)
		return c.JSON(http.StatusBadRequest, dto2.ProductNotEnoughError)
	}

	if err := h.paymentService.UpdatePaymentStatus(reqPaymentId.PaymentId, constants.SUCCESS_STATUS); err != nil {
		logger.Errorf("Failed to update payment status to success: %v", err)
		h.paymentService.UpdatePaymentStatus(reqPaymentId.PaymentId, constants.FAILED_STATUS)
		return c.JSON(http.StatusInternalServerError, dto2.GenericError)
	}
	h.productManagementService.EditProduct(&product.Name, product.Amount-payment.Amount, product.Price)
	return c.JSON(http.StatusOK, dto2.GenericSuccess)
}

func (h *productHandler) fetchPaymentAndProduct(reqPaymentId requests.PaymentIdRequest) (*entities.Payment, *entities.Product, error) {
	var payment *entities.Payment
	var product *entities.Product
	var g errgroup.Group

	g.Go(func() error {
		var err error
		payment, err = h.paymentService.FindPaymentByPaymentId(reqPaymentId)
		if err != nil {
			return &dto2.PaymentNotFound
		}
		logger.Infof("Payment: %+v", payment)
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, nil, err
	}

	g.Go(func() error {
		product = h.productManagementService.FindProductByProductId(payment.ProductID)
		if product == nil {
			return &dto2.ProductNotFound
		}
		logger.Infof("Product: %+v", product)
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, nil, err
	}

	return payment, product, nil
}
