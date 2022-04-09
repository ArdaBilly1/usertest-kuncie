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

	item := ItemInjector(config.MysqlDB)
	v1.POST("/item", item.InsertItem)
	v1.GET("/item", item.GetAllItem)
}
