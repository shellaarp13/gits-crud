package http

import (
	"myapp/entity"
	"myapp/service"

	"net/http"
	nethttp "net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateOrderBodyRequest defines all body attributes needed to add order.
type CreateOrderBodyRequest struct {
	Customer_ID   uuid.UUID `gorm:"json:"customer_id" binding:"required"`
	Customer_Name string    `gorm:"json:"customer_name" binding:"required"`
	To_street     string    `gorm:"json:"to_street" binding:"required"`
	To_city       string    `gorm:"json:"to_city" binding:"required"`
	To_zip        string    `gorm:"json:"to_zip" binding:"required"`
	Ship_date     time.Time `gorm:"json:"ship_date" binding:"required"`
}

// OrderRowResponse defines all attributes needed to fulfill for order row entity.
type OrderRowResponse struct {
	Order_Number  uuid.UUID `gorm:"json:"order_number"`
	Customer_ID   uuid.UUID `gorm:"json:"customer_id"`
	Customer_Name string    `gorm:"json:"customer_name"`
	To_street     string    `gorm:"json:"to_street"`
	To_city       string    `gorm:"json:"to_city"`
	To_zip        string    `gorm:"json:"to_zip"`
	Ship_date     time.Time `gorm:"json:"ship_date"`
}

// OrderResponse defines all attributes needed to fulfill for pic order entity.
type OrderDetailResponse struct {
	Order_Number  uuid.UUID `gorm:"json:"order_number"`
	Customer_ID   uuid.UUID `gorm:"json:"customer_id"`
	Customer_Name string    `gorm:"json:"customer_name"`
	To_street     string    `gorm:"json:"to_street"`
	To_city       string    `gorm:"json:"to_city"`
	To_zip        string    `gorm:"json:"to_zip"`
	Ship_date     time.Time `gorm:"json:"ship_date"`
}

func buildOrderRowResponse(order *entity.Order) OrderRowResponse {
	form := OrderRowResponse{
		Order_Number:  order.Order_Number,
		Customer_ID:   order.Customer_ID,
		Customer_Name: order.Customer_Name,
		To_street:     order.To_street,
		To_city:       order.To_city,
		To_zip:        order.To_zip,
		Ship_date:     order.Ship_date,
	}

	return form
}

func buildOrderDetailResponse(order *entity.Order) OrderDetailResponse {
	form := OrderDetailResponse{
		Order_Number:  order.Order_Number,
		Customer_ID:   order.Customer_ID,
		Customer_Name: order.Customer_Name,
		To_street:     order.To_street,
		To_city:       order.To_city,
		To_zip:        order.To_zip,
		Ship_date:     order.Ship_date,
	}

	return form
}

// QueryParamsOrder defines all attributes for input query params
type QueryParamsOrder struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaOrder define attributes needed for Meta
type MetaOrder struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaOrder creates an instance of Meta response.
func NewMetaOrder(limit, offset int, total int64) *MetaOrder {
	return &MetaOrder{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// OrderHandler handles HTTP request related to user flow.
type OrderHandler struct {
	service service.OrderUseCase
}

// NewOrderHandler creates an instance of OrderHandler.
func NewOrderHandler(service service.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *OrderHandler) CreateOrder(echoCtx echo.Context) error {
	var form CreateOrderBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	orderEntity := entity.NewOrder(
		uuid.Nil,
		form.Customer_ID,
		form.Customer_Name,
		form.To_street,
		form.To_city,
		form.To_zip,
		form.Ship_date,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), orderEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", orderEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *OrderHandler) GetListOrder(echoCtx echo.Context) error {
	var form QueryParamsOrder
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	order, err := handler.service.GetListOrder(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", order)
	return echoCtx.JSON(res.Status, res)

}

func (handler *OrderHandler) GetDetailOrder(echoCtx echo.Context) error {
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

	order, err := handler.service.GetDetailOrder(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", order)
	return echoCtx.JSON(res.Status, res)
}

func (handler *OrderHandler) UpdateOrder(echoCtx echo.Context) error {
	var form CreateOrderBodyRequest
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

	orderEntity := &entity.Order{
		Order_Number:  id,
		Customer_ID:   form.Customer_ID,
		Customer_Name: form.Customer_Name,
		To_street:     form.To_street,
		To_city:       form.To_city,
		To_zip:        form.To_zip,
		Ship_date:     form.Ship_date,
	}

	if err := handler.service.UpdateOrder(echoCtx.Request().Context(), orderEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *OrderHandler) DeleteOrder(echoCtx echo.Context) error {
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

	err = handler.service.DeleteOrder(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
