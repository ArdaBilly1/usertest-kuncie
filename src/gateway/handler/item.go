package handler

import (
	"net/http"
	"usertest-kuncie/src/gateway/request"
	"usertest-kuncie/src/service"
	"usertest-kuncie/utils/common"

	"github.com/labstack/echo/v4"
)

type ItemHandler interface {
	InsertItem(c echo.Context) error
	GetAllItem(c echo.Context) error
}

type itemHandler struct {
	service service.ItemService
}

func New(srv service.ItemService) ItemHandler {
	return &itemHandler{service: srv}
}

func (h *itemHandler) InsertItem(c echo.Context) error {
	req := new(request.InsertRequest)
	resp := new(common.DefaultResponse)

	if err := c.Bind(&req); err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	if err := c.Validate(req); err != nil {
		errRes := common.GenerateResponseError(err)
		return c.JSON(http.StatusBadRequest, errRes)
	}

	if err := h.service.CreateItem(req); err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp.Status = true
	resp.Message = "Success create data"

	return c.JSON(http.StatusOK, resp)
}

func (h *itemHandler) GetAllItem(c echo.Context) error {
	resp := new(common.DefaultResponse)

	data, err := h.service.GetAllItem()
	if err != nil {
		resp.Status = false
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp.Status = true
	resp.Message = "Success get data"
	resp.Data = data

	return c.JSON(http.StatusOK, resp)
}
