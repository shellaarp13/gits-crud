package http

import (
	"myapp/entity"
	"myapp/service"

	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateProductBodyRequest defines all body attributes needed to add product.
type CreateProductBodyRequest struct {
	Product_name string `json:"product_name" binding:"required"`
	Stock_P      int32  `json:"stock_p" binding:"required"`
	Product_type string `json:"product_type" binding:"required"`
	Price        int32  `json:"price" binding:"required"`
}

// ProductRowResponse defines all attributes needed to fulfill for product row entity.
type ProductRowResponse struct {
	Product_ID   uuid.UUID `json:"product_id"`
	Product_name string    `json:"product_name"`
	Stock_P      int32     `json:"stock_p"`
	Product_type string    `json:"product_type"`
	Price        int32     `json:"price"`
}

// ProductResponse defines all attributes needed to fulfill for pic product entity.
type ProductDetailResponse struct {
	Product_ID   uuid.UUID `json:"product_id"`
	Product_name string    `json:"product_name"`
	Stock_P      int32     `json:"stock_p"`
	Product_type string    `json:"product_type"`
	Price        int32     `json:"price"`
}

func buildProductRowResponse(product *entity.Product) ProductRowResponse {
	form := ProductRowResponse{
		Product_ID:   product.Product_ID,
		Product_name: product.Product_name,
		Stock_P:      product.Stock_P,
		Product_type: product.Product_type,
		Price:        product.Price,
	}
	return form
}

func buildProductDetailResponse(product *entity.Product) ProductDetailResponse {
	form := ProductDetailResponse{
		Product_ID:   product.Product_ID,
		Product_name: product.Product_name,
		Stock_P:      product.Stock_P,
		Product_type: product.Product_type,
		Price:        product.Price,
	}
	return form
}

// QueryParamsProduct defines all attributes for input query params
type QueryParamsProduct struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaProduct define attributes needed for Meta
type MetaProduct struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaProduct creates an instance of Meta response.
func NewMetaProduct(limit, offset int, total int64) *MetaProduct {
	return &MetaProduct{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// ProductHandler handles HTTP request related to user flow.
type ProductHandler struct {
	service service.ProductUseCase
}

// NewProductHandler creates an instance of ProductHandler.
func NewProductHandler(service service.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *ProductHandler) CreateProduct(echoCtx echo.Context) error {
	var form CreateProductBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	productEntity := entity.NewProduct(
		uuid.Nil,
		form.Product_name,
		form.Stock_P,
		form.Product_type,
		form.Price,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), productEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", productEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ProductHandler) GetListProduct(echoCtx echo.Context) error {
	var form QueryParamsProduct
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	product, err := handler.service.GetListProduct(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", product)
	return echoCtx.JSON(res.Status, res)

}

func (handler *ProductHandler) GetDetailProduct(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	product, err := handler.service.GetDetailProduct(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", product)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ProductHandler) UpdateProduct(echoCtx echo.Context) error {
	var form CreateProductBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	idParam := echoCtx.Param("id")

	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	_, err = handler.service.GetDetailProduct(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	productEntity := &entity.Product{
		Product_ID:   id,
		Product_name: form.Product_name,
		Stock_P:      form.Stock_P,
		Product_type: form.Product_type,
		Price:        form.Price,
	}

	if err := handler.service.UpdateProduct(echoCtx.Request().Context(), productEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ProductHandler) DeleteProduct(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	err = handler.service.DeleteProduct(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
