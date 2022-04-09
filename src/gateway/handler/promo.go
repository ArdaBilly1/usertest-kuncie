package handler

import (
	"fmt"
	"net/http"
	"usertest-kuncie/src/gateway/request"
	"usertest-kuncie/src/model"
	"usertest-kuncie/src/service"
	"usertest-kuncie/utils/common"

	"github.com/labstack/echo/v4"
)

type PromoHandler interface {
	CreatePromo(c echo.Context) error
}

type promoHandler struct {
	service service.PromoService
}

func NewPromoHandler(srv service.PromoService) PromoHandler {
	return &promoHandler{service: srv}
}

var (
	ErrorMsgPercentage error = fmt.Errorf("persentase wajib diisi")
	ErrorMsgItemBonus  error = fmt.Errorf("item bonus wajib diisi")
	ErrorMsgItemPromo  error = fmt.Errorf("item promo wajib diisi")
)

func (h *promoHandler) CreatePromo(c echo.Context) error {
	req := new(request.PromoInsertRequest)
	resp := new(common.DefaultResponse)

	if err := c.Bind(&req); err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	if err := c.Validate(req); err != nil {
		respErr := common.GenerateResponseError(err)
		return c.JSON(http.StatusBadRequest, respErr)
	}

	var promoTypeValidate interface{}
	var msgError string
	switch req.Type {
	case model.PromoTypePercentage:
		promoTypeValidate = req.Percentage
		msgError = ErrorMsgPercentage.Error()
	case model.PromoTypeBonusItem:
		promoTypeValidate = req.ItemBonusId
		msgError = ErrorMsgItemBonus.Error()
	default:
		promoTypeValidate = 0
	}

	if promoTypeValidate == 0 {
		resp.Message = msgError
		return c.JSON(http.StatusBadRequest, resp)
	} else if (promoTypeValidate == model.PromoTypeBonusItem) && (req.ItemId == 0) {
		resp.Message = ErrorMsgItemPromo.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err := h.service.CreatePromo(req); err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp.Status = true
	resp.Message = "Success insert promo"
	return c.JSON(http.StatusCreated, resp)
}
