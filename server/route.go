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
	likeSpotModel := model.NewLikeSpotModel(db)
// 	likeAreaModel := model.NewLikeAreaModel(db, likeSpotModel)

	spotPhotoModel := model.NewSpotPhotoModel()
	spotModel := model.NewSpotModel(db, spotPhotoModel, likeSpotModel)

	hotLevelModel := model.NewHotLevelModel()

	allSpotsModel := model.NewAllSpotsModel(db, hotLevelModel)
	allAreasModel := model.NewAllAreasModel(db, hotLevelModel, allSpotsModel)

	spotInfoPbView := pb.NewSpotInfoPbView(spotModel)
	allAreaInfoPbView := pb.NewAllAreasPbView(allAreasModel)
	allSpotInfoPbView := pb.NewAllSpotsPbView(allSpotsModel)

	// ルーティング
	// これはテスト用にHello, Worldしてる
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Protocol Buffersルーターによるルーティング
	pbRouter := pbrouter.NewProtocolBufferRouterByRouter(e)
	pbRouter.POST("/spotInfo/get", spotInfoPbView.GetPublishInterface(), spotInfoPbView.GetPOSTHandler)
	pbRouter.GET("/areaInfo/get/all", allAreaInfoPbView.GetGETHandler)
	pbRouter.GET("/spotInfo/get/all", allSpotInfoPbView.GetGETHandler)

	return nil
}
