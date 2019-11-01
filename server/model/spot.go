package model

import (
	"errors"
	"fmt"
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
	// DBに追加するデータ
	adding := &types.SpotInfo{
		Name:        spot.Name,
		Description: spot.Description,
		Latitude:    spot.LocatedAt.Latitude,
		Longitude:   spot.LocatedAt.Longitude,
		StartingAt:  spot.StartingAt,
		EndingAt:    spot.EndingAt,
		AreaInfoID:  spot.AreaID,
	}

	// 画像の保存
	photoFileName, err := model.spotPhotoModel.Add(spot.PhotoReader)
	if err != nil {
		return err
	}
	adding.PhotoFileName = photoFileName

	// DBに追加
	result := model.db.Create(adding)
	err = result.Error
	if err != nil {
		errMsg := "SpotModel.Add(): " + err.Error()
		return errors.New(errMsg)
	}
	affected := result.RowsAffected
	if affected <= 0 {
		errMsg := fmt.Sprintf("SpotModel.Add(): Spot(%s) can't be added.", adding.Name)
		return errors.New(errMsg)
	}

	return nil
}

func (model *SpotModelImpl) Like(spotID uint) (err error) {
	panic("implement me")
}
