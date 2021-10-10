package http

import (
	"fmt"
	"myapp/entity"
	"myapp/service"

	"net/http"
	nethttp "net/http"

	"myapp/internal/tool/encrypter"

	"github.com/labstack/echo/v4"
)

// CreateAccountBodyRequest defines all body attributes needed to add account.
type CreateAccountBodyRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AccountRowResponse defines all attributes needed to fulfill for account row entity.
type AccountRowResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AccountResponse defines all attributes needed to fulfill for pic account entity.
type AccountDetailResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func buildAccountRowResponse(account *entity.Account) AccountRowResponse {
	form := AccountRowResponse{
		Username: account.Username,
		Password: account.Password,
	}

	return form
}

func buildAccountDetailResponse(account *entity.Account) AccountDetailResponse {
	form := AccountDetailResponse{
		Username: account.Username,
		Password: account.Password,
	}

	return form
}

// QueryParamsAccount defines all attributes for input query params
type QueryParamsAccount struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaAccount define attributes needed for Meta
type MetaAccount struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaAccount creates an instance of Meta response.
func NewMetaAccount(limit, offset int, total int64) *MetaAccount {
	return &MetaAccount{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// AccountHandler handles HTTP request related to user flow.
type AccountHandler struct {
	service service.AccountUseCase
}

// NewAccountHandler creates an instance of AccountHandler.
func NewAccountHandler(service service.AccountUseCase) *AccountHandler {
	return &AccountHandler{
		service: service,
	}
}

// Create handles article creation.
// It will reject the request if the request doesn't have required data,
func (handler *AccountHandler) CreateAccount(echoCtx echo.Context) error {
	var form CreateAccountBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	password, err := encrypter.NewBcrypt().Encrypt(form.Password)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	accountEntity := entity.NewAccount(
		form.Username,
		password,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), accountEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", accountEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *AccountHandler) GetListAccount(echoCtx echo.Context) error {
	var form QueryParamsAccount
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	account, err := handler.service.GetListAccount(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", account)
	return echoCtx.JSON(res.Status, res)

}

func (handler *AccountHandler) GetDetailAccount(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	fmt.Println(idParam)
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	account, err := handler.service.GetDetailAccount(echoCtx.Request().Context(), idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", account)
	return echoCtx.JSON(res.Status, res)
}

func (handler *AccountHandler) UpdateAccount(echoCtx echo.Context) error {
	var form CreateAccountBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	idParam := echoCtx.Param("id")

	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	_, err := handler.service.GetDetailAccount(echoCtx.Request().Context(), idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	password, err := encrypter.NewBcrypt().Encrypt(form.Password)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	accountEntity := &entity.Account{
		Username: form.Username,
		Password: password,
	}

	if err := handler.service.UpdateAccount(echoCtx.Request().Context(), idParam, accountEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *AccountHandler) DeleteAccount(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	err := handler.service.DeleteAccount(echoCtx.Request().Context(), idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
