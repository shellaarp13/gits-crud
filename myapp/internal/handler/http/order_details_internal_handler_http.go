package http

import (
	"myapp/entity"
	"myapp/service"

	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateOrderDetailsBodyRequest defines all body attributes needed to add order.
type CreateOrderDetailsBodyRequest struct {
	Order_details_id uuid.UUID `gorm:"json:"customer_id" binding:"required"`
	Order_number      string    `gorm:"json:"customer_name" binding:"required"`
	Product_id       string    `gorm:"json:"to_street" binding:"required"`
	Quantity_product string    `gorm:"json:"to_city" binding:"required"`
}

// OrderDetailsRowResponse defines all attributes needed to fulfill for orderdetails row entity.
type OrderDetailsRowResponse struct {
	Order_details_id uuid.UUID `gorm:"json:"order_details_id"`
	Order_Number     string    `gorm:"json:"order_number"`
	Product_id       string    `gorm:"json:"product_id"`
	Quantity_product string    `gorm:"json:"quantity_product"`
}

// OrderDetailsResponse defines all attributes needed to fulfill for pic orderdetails entity.
type OrderDetailsDetailResponse struct {
	Order_details_id uuid.UUID `gorm:"json:"order_details_id"`
	Order_Number     string    `gorm:"json:"order_number"`
	Product_id       string    `gorm:"json:"product_id"`
	Quantity_product string    `gorm:"json:"quantity_product"`
}

func buildOrderDetailsRowResponse(order_details *entity.OrderDetails) OrderDetailsRowResponse {
	form := OrderDetailsRowResponse{
		Order_details_id: order_details.Order_details_id,
		Order_Number:     order_details.Order_Number,
		Product_id:       order_details.Product_id,
		Quantity_product: order_details.Quantity_product,
	}

	return form
}

func buildOrderDetailsDetailResponse(order_details *entity.OrderDetails) OrderDetailsDetailResponse {
	form := OrderDetailsDetailResponse{
		Order_details_id: order_details.oder_details_id,
		Order_Number:     order_details.Order_Number,
		Product_id:       order_details.Product_id,
		Quantity_product: order_details.Quantity_product,
	}

	return form
}

// QueryParamsOrderDetails defines all attributes for input query params
type QueryParamsOrderDetails struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaOrderDetails define attributes needed for Meta
type MetaOrderDetails struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaOrderDetails creates an instance of Meta response.
func NewMetaOrderDetails(limit, offset int, total int64) *MetaOrderDetails {
	return &MetaOrderDetails{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// OrderDetailsHandler handles HTTP request related to user flow.
type OrderDetailsHandler struct {
	service service.OrderUseCase
}

// NewOrderDetailsHandler creates an instance of OrderDetailsHandler.
func NewOrderDetailsHandler(service service.OrderDetailsUseCase) *OrderDetailsHandler {
	return &OrderDetailsHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *OrderDetailsHandler) CreateOrderDetails(echoCtx echo.Context) error {
	var form CreateOrderDetailsBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	order_detailsEntity := entity.NewOrderDetails(
		uuid.Nil,
		form.Oder_number,
		form.Product_id,
		form.Quantity_product,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), order_detailsEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", order_detailsEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *OrderDetailsHandler) GetListOrderDetails(echoCtx echo.Context) error {
	var form QueryParamsOrderDetails
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	order_details, err := handler.service.GetListOrderDetails(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", order_details)
	return echoCtx.JSON(res.Status, res)

}

func (handler *OrderDetailsHandler) GetDetailOrderDetails(echoCtx echo.Context) error {
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

	order_details, err := handler.service.GetDetailOrderDetails(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", order_details)
	return echoCtx.JSON(res.Status, res)
}

func (handler *OrderDetailsHandler) UpdateOrderDetails(echoCtx echo.Context) error {
	var form CreateOrderDetailsBodyRequest
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

	_, err = handler.service.GetDetailOrder(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	order_detailsEntity := &entity.OrderDetails{
		Order_details_id: id,
		Order_number:     form.Oder_number,
		Product_id:       form.Product_id,
		Quantity_product: form.Quantity_product,
	}

	if err := handler.service.UpdateOrderDetails(echoCtx.Request().Context(), order_detailsEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *OrderDetailsHandler) DeleteOrderDetails(echoCtx echo.Context) error {
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

	err = handler.service.DeleteOrderDetails(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
