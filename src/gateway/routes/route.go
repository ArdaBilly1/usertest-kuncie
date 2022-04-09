package routes

import (
	"net/http"
	"usertest-kuncie/utils/config"

	"github.com/labstack/echo/v4"
)

func V1(e *echo.Echo) {
	v1 := e.Group("v1")

	v1.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "echo says : 'im fine'")
	})

	itemHandler := ItemInjector(config.MysqlDB)
	item := v1.Group("/item")
	item.POST("", itemHandler.InsertItem)
	item.GET("", itemHandler.GetAllItem)

	promoHandler := PromoInjector(config.MysqlDB)
	promo := v1.Group("/promo")
	promo.POST("", promoHandler.CreatePromo)

	orderHandler := OrderInjectory(config.MysqlDB)
	order := v1.Group("/order")
	order.POST("/checkout", orderHandler.Checkout)
}
