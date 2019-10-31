package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
)

func Route(e *echo.Echo, db *gorm.DB) error {
	// ControllerやModelの準備


	// ルーティング
	// これはテスト用にHello, Worldしてる
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return nil
}
