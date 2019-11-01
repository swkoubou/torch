package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/swkoubou/torch/server/model/types"
)

type SpotModelImpl struct {
	db             *gorm.DB
	spotPhotoModel SpotPhotoModel
}

func NewSpotModel(db *gorm.DB, spotPhotoModel SpotPhotoModel) SpotModel {
	return &SpotModelImpl{db: db, spotPhotoModel: spotPhotoModel}
}

func (model *SpotModelImpl) Get(spotID uint) (spot *types.SpotInfo, err error) {
	spot = &types.SpotInfo{}
	result := model.db.First(&spot, spotID)

	err = result.Error
	if result.RecordNotFound() {
		return nil, errors.New("not found")
	}
	if err != nil {
		return nil, err
	}

	return spot, nil
}

func (model *SpotModelImpl) Add(spot *types.NewSpotInfo) (err error) {
	panic("implement me")
}

func (model *SpotModelImpl) Like(spotID uint) (err error) {
	panic("implement me")
}
