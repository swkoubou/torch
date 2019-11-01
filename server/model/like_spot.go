package model

import (
	"github.com/jinzhu/gorm"
	"github.com/swkoubou/torch/server/model/types"
)

type LikeSpotModelImpl struct {
	db *gorm.DB
}

func (model *LikeSpotModelImpl) Like(spotID uint) (count uint, err error) {
	panic("implement me")
}

func (model *LikeSpotModelImpl) CountLikes(spotID uint) (count uint, err error) {
	panic("implement me")
}

func (model *LikeSpotModelImpl) CountAllLikes(targetSpots []types.SpotInfo) (counted []types.SpotInfo, err error) {
	panic("implement me")
}
