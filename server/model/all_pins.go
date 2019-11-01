package model

import (
	"github.com/jinzhu/gorm"
	"github.com/swkoubou/torch/server/model/types"
)

type AllSpotsModelImpl struct {
	db            *gorm.DB
	hotLevelModel HotLevelModel
}

func NewAllSpotsModel(db *gorm.DB, hotLevelModel HotLevelModel) AllSpotsModel {
	return &AllSpotsModelImpl{
		db:            db,
		hotLevelModel: hotLevelModel,
	}
}

func (model *AllSpotsModelImpl) Get() (spots []types.SpotInfo, err error) {
	panic("implement me")
}

func (model *AllSpotsModelImpl) GetWithHotLevel() (spots []types.SpotInfo, err error) {
	panic("implement me")
}
