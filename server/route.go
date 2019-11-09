package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/swkoubou/torch/server/lib/pbrouter"
	"github.com/swkoubou/torch/server/model"
	"github.com/swkoubou/torch/server/view/pb"
	"github.com/swkoubou/torch/server/view/rest"
)

func Route(e *echo.Echo, db *gorm.DB) error {
	// ViewやModelの準備
	likeSpotModel := model.NewLikeSpotModel(db)
	likeAreaModel := model.NewLikeAreaModel(db, likeSpotModel)

	spotPhotoModel := model.NewSpotPhotoModel()
	spotModel := model.NewSpotModel(db, spotPhotoModel, likeSpotModel)

	hotLevelModel := model.NewHotLevelModel()

	allSpotsModel := model.NewAllSpotsModel(db, hotLevelModel, likeSpotModel)
	allAreasModel := model.NewAllAreasModel(db, hotLevelModel, allSpotsModel, likeAreaModel)

	spotInfoView := rest.NewSpotInfoRESTView(spotModel)
	spotInfoPbView := pb.NewSpotInfoPbView(spotModel)
	allAreaInfoPbView := pb.NewAllAreasPbView(allAreasModel)
	allSpotInfoPbView := pb.NewAllSpotsPbView(allSpotsModel)
	likeSpotPbView := pb.NewLikeSpotPbView(likeSpotModel)
	likeAreaPbView := pb.NewLikeAreaPbView(likeAreaModel)

	// ルーティング
	e.POST("/spotInfo/create", spotInfoView.GetPUTHandler)

	pbRouter := pbrouter.NewProtocolBufferRouterByRouter(e)
	pbRouter.POST("/spotInfo/get", spotInfoPbView.GetPublishInterface(), spotInfoPbView.GetPOSTHandler)
	pbRouter.GET("/spotInfo/get/all", allSpotInfoPbView.GetGETHandler)
	pbRouter.POST("/spotInfo/like", likeSpotPbView.GetPublishInterface(), likeSpotPbView.GetPOSTHandler)

	pbRouter.GET("/areaInfo/get/all", allAreaInfoPbView.GetGETHandler)
	pbRouter.POST("/areaInfo/like", likeAreaPbView.GetPublishInterface(), likeAreaPbView.GetPOSTHandler)

	return nil
}
