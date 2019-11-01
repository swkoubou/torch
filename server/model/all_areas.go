package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/swkoubou/torch/server/model/types"
)

type AllAreasModelImpl struct {
	db            *gorm.DB
	hotLevelModel HotLevelModel
	allSpotsModel AllSpotsModel
}

func NewAllAreasModel(db *gorm.DB, hotLevelModel HotLevelModel, allSpotsModel AllSpotsModel) AllAreasModel {
	return &AllAreasModelImpl{
		db:            db,
		hotLevelModel: hotLevelModel,
		allSpotsModel: allSpotsModel,
	}
}

func (model *AllAreasModelImpl) Get() (areas []types.AreaInfo, err error) {
	// スポット情報の取得
	areas = []types.AreaInfo{}

	result := model.db.Find(&areas)
	err = result.Error
	if err != nil {
		return nil, errors.New("AllAreasModel.Get(): " + err.Error())
	}

	// RecordNotFoundはFind系の関数では出ないエラーなので別途カウントする
	// https://github.com/jinzhu/gorm/issues/228
	if len(areas) <= 0 {
		errMsg := "AllAreasModel.Get(): No rows detected."
		return nil, errors.New(errMsg)

	}

	return areas, nil
}

func (model *AllAreasModelImpl) GetWithHotLevel() (areas []types.AreaInfo, err error) {
	// 上の関数では盛り上がり具合をロードしないので、
	// この関数では盛り上がり具合も計算するようにする

	// エリア情報の取得
	areas, err = model.Get()
	if err != nil {
		return nil, err
	}

	// スポット情報の取得
	spots, err := model.allSpotsModel.Get()

	// 盛り上がり具合をロードする
	areas, err = model.hotLevelModel.CalcAreaHotLevel(areas, spots)
	if err != nil {
		return nil, err
	}

	return areas, nil
}
