package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/swkoubou/torch/server/model/types"
)

type AllSpotsModelImpl struct {
	db            *gorm.DB
	hotLevelModel HotLevelModel
	likeSpotModel LikeSpotModel
}

func NewAllSpotsModel(
	db *gorm.DB,
	hotLevelModel HotLevelModel,
	likeSpotModel LikeSpotModel,
) AllSpotsModel {
	return &AllSpotsModelImpl{
		db:            db,
		hotLevelModel: hotLevelModel,
		likeSpotModel: likeSpotModel,
	}
}

func (model *AllSpotsModelImpl) Get() (spots []types.SpotInfo, err error) {
	// スポット情報の取得
	spots = []types.SpotInfo{}

	result := model.db.Preload("AreaInfo").Find(&spots)
	err = result.Error
	if err != nil {
		return nil, errors.New("AllSpotsModel.Get(): " + err.Error())
	}

	// RecordNotFoundはFind系の関数では出ないエラーなので別途カウントする
	// https://github.com/jinzhu/gorm/issues/228
	if len(spots) <= 0 {
		errMsg := "AllSpotsModel.Get(): No rows detected."
		return nil, errors.New(errMsg)
	}

	countedSpots, err := model.likeSpotModel.CountAllLikes(spots)
	if err != nil {
		errMsg := "AllAreasModel.Get(): Can't count area likes.; " + err.Error()
		return nil, errors.New(errMsg)
	}

	return countedSpots, nil
}

func (model *AllSpotsModelImpl) GetWithHotLevel() (spots []types.SpotInfo, err error) {
	// 上の関数では盛り上がり具合をロードしないので、
	// この関数では盛り上がり具合も計算するようにする

	// スポット情報の取得
	spots, err = model.Get()
	if err != nil {
		return nil, err
	}

	// 盛り上がり具合をロードする
	spotsWithHot, err := model.hotLevelModel.CalcSpotHotLevel(spots)
	if err != nil {
		return nil, err
	}

	return spotsWithHot, nil
}
