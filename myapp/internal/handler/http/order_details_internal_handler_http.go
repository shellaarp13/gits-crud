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
	Order_Number     uuid.UUID `gorm:"json:"order_details_name" binding:"required"`
	Product_ID       uuid.UUID `gorm:"json:"to_street" binding:"required"`
	Quantity_product int32     `gorm:"json:"to_city" binding:"required"`
}

// OrderDetailsRowResponse defines all attributes needed to fulfill for orderdetails row entity.
type OrderDetailsRowResponse struct {
	Order_details_id uuid.UUID `gorm:"json:"order_details_id"`
	Order_Number     uuid.UUID `gorm:"json:"order_number"`
	Product_ID       uuid.UUID `gorm:"json:"product_id"`
	Quantity_product int32     `gorm:"json:"quantity_product"`
}

// OrderDetailsResponse defines all attributes needed to fulfill for pic orderdetails entity.
type OrderDetailsDetailResponse struct {
	Order_details_id uuid.UUID `gorm:"json:"order_details_id"`
	Order_Number     uuid.UUID `gorm:"json:"order_number"`
	Product_ID       uuid.UUID `gorm:"json:"product_id"`
	Quantity_product int32     `gorm:"json:"quantity_product"`
}

func buildOrderDetailsRowResponse(order_details *entity.Order_Details) OrderDetailsRowResponse {
	form := OrderDetailsRowResponse{
		Order_details_id: order_details.Order_details_id,
		Order_Number:     order_details.Order_Number,
		Product_ID:       order_details.Product_ID,
		Quantity_product: order_details.Quantity_product,
	}

	return form
}

func buildOrderDetailsDetailResponse(order_details *entity.Order_Details) OrderDetailsDetailResponse {
	form := OrderDetailsDetailResponse{
		Order_details_id: order_details.Order_details_id,
		Order_Number:     order_details.Order_Number,
		Product_ID:       order_details.Product_ID,
		Quantity_product: order_details.Quantity_product,
	}

	return form
}

// QueryParamsOrder_Details defines all attributes for input query params
type QueryParamsOrder_Details struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaOrder_Details define attributes needed for Meta
type MetaOrder_Details struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaOrder_Details creates an instance of Meta response.
func NewMetaOrder_Details(limit, offset int, total int64) *MetaOrder_Details {
	return &MetaOrder_Details{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// Order_DetailsHandler handles HTTP request related to user flow.
type Order_DetailsHandler struct {
	service service.Order_DetailsUseCase
}

// NewOrder_DetailsHandler creates an instance of Order_DetailsHandler.
func NewOrder_DetailsHandler(service service.Order_DetailsUseCase) *Order_DetailsHandler {
	return &Order_DetailsHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *Order_DetailsHandler) CreateOrder_Details(echoCtx echo.Context) error {
	var form CreateOrderDetailsBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	order_detailsEntity := entity.NewOrderDetails(
		uuid.Nil,
		form.Order_Number,
		form.Product_ID,
		form.Quantity_product,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), order_detailsEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", order_detailsEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *Order_DetailsHandler) GetListOrder_Details(echoCtx echo.Context) error {
	var form QueryParamsOrder_Details
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	order_details, err := handler.service.GetListOrder_Details(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", order_details)
	return echoCtx.JSON(res.Status, res)

}

func (handler *Order_DetailsHandler) GetDetailOrder_Details(echoCtx echo.Context) error {
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

	order_details, err := handler.service.GetDetailOrder_Details(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", order_details)
	return echoCtx.JSON(res.Status, res)
}

func (handler *Order_DetailsHandler) UpdateOrder_Details(echoCtx echo.Context) error {
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

	_, err = handler.service.GetDetailOrder_Details(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	order_detailsEntity := &entity.Order_Details{
		Order_details_id: id,
		Order_Number:     form.Order_Number,
		Product_ID:       form.Product_ID,
		Quantity_product: form.Quantity_product,
	}

	if err := handler.service.UpdateOrder_Details(echoCtx.Request().Context(), order_detailsEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *Order_DetailsHandler) DeleteOrder_Details(echoCtx echo.Context) error {
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

	err = handler.service.DeleteOrder_Details(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
