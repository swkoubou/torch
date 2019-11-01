package model

import (
	"github.com/jinzhu/gorm"
	"github.com/swkoubou/torch/server/model/types"
)

type AllPinsModelImpl struct {
	db            *gorm.DB
	hotLevelModel HotLevelModel
}

func NewAllPinsModel(db *gorm.DB, hotLevelModel HotLevelModel) AllPinsModel {
	return &AllPinsModelImpl{
		db:            db,
		hotLevelModel: hotLevelModel,
	}
}

func (model *AllPinsModelImpl) Get() (areas []types.PinInfo, err error) {
	panic("implement me")
}

func (model *AllPinsModelImpl) GetWithHotLevel() (areas []types.PinInfo, err error) {
	panic("implement me")
}
