package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/swkoubou/torch/server/model/types"
	"time"
)

type LikeSpotModelImpl struct {
	db *gorm.DB
}

func NewLikeSpotModel(db *gorm.DB) LikeSpotModel {
	return &LikeSpotModelImpl{db: db}
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
	littleBitBefore := time.Now().Add(-1 * time.Minute * 30)
	result := model.db.Model(&types.LikeSpot{}).Where("spot_info_id = ?", spotID).Where("created_at > ?", littleBitBefore).Count(&count)
	err = result.Error
	if err != nil {
		errMsg := "LikeSpotModel.CountLikes(): " + err.Error()
		return 0, errors.New(errMsg)
	}

	return count, nil
}

func (model *LikeSpotModelImpl) CountAllLikes(targetSpots []types.SpotInfo) (counted []types.SpotInfo, err error) {
	var countedSpots []types.SpotInfo

	for _, v := range targetSpots {
		v.Likes, err = model.CountLikes(v.Model.ID)
		if err != nil {
			return nil, err
		}

		countedSpots = append(countedSpots, v)
	}

	return countedSpots, nil
}
