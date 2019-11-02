package model

import (
	"errors"
	"fmt"
	. "github.com/ahmetb/go-linq"
	"github.com/jinzhu/gorm"
	"github.com/swkoubou/torch/server/model/types"
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
	result := model.db.Where("id = ?", spotID).Count(&count)
	err = result.Error
	if err != nil {
		errMsg := "LikeSpotModel.CountLikes(): " + err.Error()
		return 0, errors.New(errMsg)
	}

	return count, nil
}

func (model *LikeSpotModelImpl) CountAllLikes(targetSpots []types.SpotInfo) (counted []types.SpotInfo, err error) {
	From(targetSpots).
		Select(func(a interface{}) interface{} {
			spot, isCast := a.(types.SpotInfo)
			if !isCast {
				return a // キャストできなくてpanicは避けたいのでそのまま返す
			}

			spot.Likes, err = model.CountLikes(spot.ID)
			if err != nil {
				spot.Likes = 0
				return spot
			}

			return spot
		}).ToSlice(counted)

	return counted, nil
}
