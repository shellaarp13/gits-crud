package http

import (
	"myapp/entity"
	"myapp/service"

	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateCustomerBodyRequest defines all body attributes needed to add customer.
type CreateCustomerBodyRequest struct {
	First_Name string `json:"first_name" binding:"required"`
	Last_Name  string `json:"last_name" binding:"required"`
	Street     string `json:"street" binding:"required"`
	Zip        string `json:"zip" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
}

// CustomerRowResponse defines all attributes needed to fulfill for customer row entity.
type CustomerRowResponse struct {
	Customer_ID uuid.UUID `json:"customer_id"`
	First_Name  string    `json:"first_name"`
	Last_Name   string    `json:"last_name"`
	Street      string    `json:"street"`
	Zip         string    `json:"zip"`
	Phone       string    `json:"phone"`
}

// CustomerResponse defines all attributes needed to fulfill for pic customer entity.
type CustomerDetailResponse struct {
	Customer_ID uuid.UUID `json:"customer_id"`
	First_Name  string    `json:"first_name"`
	Last_Name   string    `json:"last_name"`
	Street      string    `json:"street"`
	Zip         string    `json:"zip"`
	Phone       string    `json:"phone"`
}

func buildCustomerRowResponse(customer *entity.Customer) CustomerRowResponse {
	form := CustomerRowResponse{
		Customer_ID: customer.Customer_ID,
		First_Name:  customer.First_Name,
		Last_Name:   customer.Last_Name,
		Street:      customer.Street,
		Zip:         customer.Zip,
		Phone:       customer.Phone,
	}

	return form
}

func buildCustomerDetailResponse(customer *entity.Customer) CustomerDetailResponse {
	form := CustomerDetailResponse{
		Customer_ID: customer.Customer_ID,
		First_Name:  customer.First_Name,
		Last_Name:   customer.Last_Name,
		Street:      customer.Street,
		Zip:         customer.Zip,
		Phone:       customer.Phone,
	}

	return form
}

// QueryParamsCustomer defines all attributes for input query params
type QueryParamsCustomer struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaCustomer define attributes needed for Meta
type MetaCustomer struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaCustomer creates an instance of Meta response.
func NewMetaCustomer(limit, offset int, total int64) *MetaCustomer {
	return &MetaCustomer{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// CustomerHandler handles HTTP request related to user flow.
type CustomerHandler struct {
	service service.CustomerUseCase
}

// NewCustomerHandler creates an instance of CustomerHandler.
func NewCustomerHandler(service service.CustomerUseCase) *CustomerHandler {
	return &CustomerHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *CustomerHandler) CreateCustomer(echoCtx echo.Context) error {
	var form CreateCustomerBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	customerEntity := entity.NewCustomer(
		uuid.Nil,
		form.First_Name,
		form.Last_Name,
		form.Street,
		form.Zip,
		form.Phone,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), customerEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", customerEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *CustomerHandler) GetListCustomer(echoCtx echo.Context) error {
	var form QueryParamsCustomer
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	customer, err := handler.service.GetListCustomer(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", customer)
	return echoCtx.JSON(res.Status, res)

}

func (handler *CustomerHandler) GetDetailCustomer(echoCtx echo.Context) error {
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

	customer, err := handler.service.GetDetailCustomer(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", customer)
	return echoCtx.JSON(res.Status, res)
}

func (handler *CustomerHandler) UpdateCustomer(echoCtx echo.Context) error {
	var form CreateCustomerBodyRequest
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

	_, err = handler.service.GetDetailCustomer(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	customerEntity := &entity.Customer{
		Customer_ID: id,
		First_Name:  form.First_Name,
		Last_Name:   form.Last_Name,
		Street:      form.Street,
		Zip:         form.Zip,
		Phone:       form.Phone,
	}

	if err := handler.service.UpdateCustomer(echoCtx.Request().Context(), customerEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *CustomerHandler) DeleteCustomer(echoCtx echo.Context) error {
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

	err = handler.service.DeleteCustomer(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
