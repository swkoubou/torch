package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/swkoubou/torch/server/lib/pbrouter"
	"github.com/swkoubou/torch/server/model"
	"github.com/swkoubou/torch/server/view/pb"
	"net/http"
)

func Route(e *echo.Echo, db *gorm.DB) error {
	// ViewやModelの準備
	spotPhotoModel := model.NewSpotPhotoModel()
	spotModel := model.NewSpotModel(db, spotPhotoModel)

	spotInfoPbView := pb.NewSpotInfoPbView(spotModel)

	// ルーティング
	// これはテスト用にHello, Worldしてる
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Protocol Buffersルーターによるルーティング
	pbRouter := pbrouter.NewProtocolBufferRouterByRouter(e)
	pbRouter.POST("/spot/info", spotInfoPbView.GetPublishInterface(), spotInfoPbView.GetPOSTHandler)

	return nil
}
