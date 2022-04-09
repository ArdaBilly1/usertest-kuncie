package handler

import (
	"net/http"
	"usertest-kuncie/src/gateway/request"
	"usertest-kuncie/src/service"
	"usertest-kuncie/utils/common"

	"github.com/labstack/echo/v4"
)

type OrderHandlerContract interface {
	Checkout(c echo.Context) error
}

type orderHandler struct {
	service service.OrderService
}

func NewOrderHandler(srv service.OrderService) OrderHandlerContract {
	return &orderHandler{service: srv}
}

func (h *orderHandler) Checkout(c echo.Context) error {
	req := new(request.OrderRequest)
	resp := new(common.DefaultResponse)

	if err := c.Bind(&req); err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	if err := c.Validate(req); err != nil {
		respErr := common.GenerateResponseError(err)
		return c.JSON(http.StatusBadRequest, respErr)
	}

	if err := h.service.Checkout(req); err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp.Status = true
	resp.Message = "success add items to cart"
	return c.JSON(http.StatusCreated, resp)
}
