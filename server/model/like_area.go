package model

import (
	"errors"
	"fmt"
	. "github.com/ahmetb/go-linq"
	"github.com/jinzhu/gorm"
	"github.com/swkoubou/torch/server/model/types"
)

type LikeAreaModelImpl struct {
	db            *gorm.DB
	likeSpotModel LikeSpotModel
}

func NewLikeAreaModel(db *gorm.DB, likeSpotModel LikeSpotModel) LikeAreaModel {
	return &LikeAreaModelImpl{
		db:            db,
		likeSpotModel: likeSpotModel,
	}
}

func (model *LikeAreaModelImpl) Like(areaID uint) (err error) {
	adding := &types.LikeArea{
		AreaInfoID: areaID,
	}

	result := model.db.Create(&adding)
	err = result.Error
	if err != nil {
		errMsg := "LikeAreaModel.Like(): " + err.Error()
		return errors.New(errMsg)
	}
	affected := result.RowsAffected
	if affected <= 0 {
		errMsg := fmt.Sprintf("LikeAreaModel.Like(): LikeArea(AreaID: %s) can't be added.", adding.AreaInfoID)
		return errors.New(errMsg)
	}

	return nil
}

func (model *LikeAreaModelImpl) CountLikes(areaID uint) (count uint, err error) {
	result := model.db.Where("id = ?", areaID).Count(&count)
	err = result.Error
	if err != nil {
		errMsg := "LikeAreaModel.CountLikes(): " + err.Error()
		return 0, errors.New(errMsg)
	}

	return count, nil
}

func (model *LikeAreaModelImpl) CountAllLikes(targetAreas []types.AreaInfo, targetSpots []types.SpotInfo) (counted []types.AreaInfo, err error) {
	From(targetAreas).
		Select(func(a interface{}) interface{} {
			// エリアのキャスト
			area, isCast := a.(types.AreaInfo)
			if !isCast {
				return a // キャストできなくてpanicは避けたいのでそのまま返す
			}

			// エリアのいいね数を取得
			area.Likes, err = model.CountLikes(area.ID)
			if err != nil {
				area.Likes = 0
				return area
			}

			// エリア内のすべてのスポットを取得
			var locatedSpots []types.SpotInfo
			From(targetSpots).
				Where(func(s interface{}) bool {
					spot, isCast := s.(types.SpotInfo)
					if !isCast {
						return false
					}

					return spot.Model.ID == area.ID
				}).
				ToSlice(&locatedSpots)

			// スポットのいいねをとりあえず全部取得
			locatedSpots, err = model.likeSpotModel.CountAllLikes(locatedSpots)
			if err != nil {
				area.Likes = 0
				return area
			}

			// 全部たす
			for _, v := range locatedSpots {
				area.Likes += v.Likes
			}

			return area
		}).ToSlice(counted)

	return counted, nil
}
