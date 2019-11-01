package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/swkoubou/torch/server/model/types"
)

type LikeSpotModelImpl struct {
	db *gorm.DB
}

func (model *LikeSpotModelImpl) Like(spotID uint) (err error) {
	adding := &types.LikeSpot{
		SpotInfoID: spotID,
	}

	result := model.db.Create(&adding)
	err = result.Error
	if err != nil {
		errMsg := "LikeSpotModel.Like(): " + err.Error()
		return errors.New(errMsg)
	}
	affected := result.RowsAffected
	if affected <= 0 {
		errMsg := fmt.Sprintf("LikeSpotModel.Like(): LikeSpot(SpotID: %s) can't be added.", adding.SpotInfoID)
		return errors.New(errMsg)
	}

	return nil
}

func (model *LikeSpotModelImpl) CountLikes(spotID uint) (count uint, err error) {
	panic("implement me")
}

func (model *LikeSpotModelImpl) CountAllLikes(targetSpots []types.SpotInfo) (counted []types.SpotInfo, err error) {
	panic("implement me")
}
