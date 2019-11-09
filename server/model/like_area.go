package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/swkoubou/torch/server/model/types"
	"time"
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
	// いいねした履歴を保存する
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
	// TODO : ここでは十把一絡げに30分前からのいいねを取得しているが、全期間のいいね数をとりたいケースもあるので、なんとかする

	// ちょっと前(30分前)
	littleBitBefore := time.Now().Add(-1 * time.Minute * 30)

	// ちょっと前から現在までを検索対象にいいね数を取得する
	result := model.db.Model(&types.LikeArea{}).Where("area_info_id = ?", areaID).Where("created_at > ?", littleBitBefore).Count(&count)
	err = result.Error
	if err != nil {
		errMsg := "LikeAreaModel.CountLikes(): " + err.Error()
		return 0, errors.New(errMsg)
	}

	return count, nil
}

func (model *LikeAreaModelImpl) CountAllLikes(targetAreas []types.AreaInfo, targetSpots []types.SpotInfo) (counted []types.AreaInfo, err error) {
	var countedAreas []types.AreaInfo

	for _, area := range targetAreas {
		// エリアのいいね数を取得
		area.Likes, err = model.CountLikes(area.ID)
		if err != nil {
			return nil, err
		}

		// エリア内のすべてのスポットを取得
		for _, spot := range targetSpots {
			if spot.AreaInfoID != area.Model.ID {
				continue
			}

			area.Likes += spot.Likes
		}

		countedAreas = append(countedAreas, area)
	}

	return countedAreas, nil
}
