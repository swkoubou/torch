package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/swkoubou/torch/server/lib/pbrouter"
	"github.com/swkoubou/torch/server/model"
	"github.com/swkoubou/torch/server/view/pb"
	"github.com/swkoubou/torch/server/view/rest"
	"net/http"
)

func Route(e *echo.Echo, db *gorm.DB) error {
	// ViewやModelの準備
	likeSpotModel := model.NewLikeSpotModel(db)
	likeAreaModel := model.NewLikeAreaModel(db, likeSpotModel)

	spotPhotoModel := model.NewSpotPhotoModel()
	spotModel := model.NewSpotModel(db, spotPhotoModel, likeSpotModel)

	hotLevelModel := model.NewHotLevelModel()

	allSpotsModel := model.NewAllSpotsModel(db, hotLevelModel)
	allAreasModel := model.NewAllAreasModel(db, hotLevelModel, allSpotsModel)

	spotInfoView := rest.NewSpotInfoRESTView(spotModel)
	spotInfoPbView := pb.NewSpotInfoPbView(spotModel)
	allAreaInfoPbView := pb.NewAllAreasPbView(allAreasModel)
	allSpotInfoPbView := pb.NewAllSpotsPbView(allSpotsModel)
	likeSpotPbView := pb.NewLikeSpotPbView(likeSpotModel)
	likeAreaPbView := pb.NewLikeAreaPbView(likeAreaModel)

	// ルーティング
	// これはテスト用にHello, Worldしてる
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/spotInfo/create", spotInfoView.GetPUTHandler)

	// Protocol Buffersルーターによるルーティング
	pbRouter := pbrouter.NewProtocolBufferRouterByRouter(e)
	pbRouter.POST("/spotInfo/get", spotInfoPbView.GetPublishInterface(), spotInfoPbView.GetPOSTHandler)

	pbRouter.GET("/areaInfo/get/all", allAreaInfoPbView.GetGETHandler)
	pbRouter.GET("/spotInfo/get/all", allSpotInfoPbView.GetGETHandler)

	pbRouter.POST("/areaInfo/like", likeAreaPbView.GetPublishInterface(), likeAreaPbView.GetPOSTHandler)
	pbRouter.POST("/spotInfo/like", likeSpotPbView.GetPublishInterface(), likeSpotPbView.GetPOSTHandler)

	return nil
}
