package model

import (
	. "github.com/ahmetb/go-linq"
	"github.com/swkoubou/torch/server/model/types"
)

type HotLevelModelImpl struct{}

func NewHotLevelModel() HotLevelModel {
	return &HotLevelModelImpl{}
}

func (model *HotLevelModelImpl) CalcAreaHotLevel(targetAreas []types.AreaInfo, targetSpots []types.SpotInfo) (calculated []types.AreaInfo, err error) {
	// TODO : いつか計算する
	From(targetAreas).
		Select(func(a interface{}) interface{} {
			area, isCast := a.(types.AreaInfo)
			if !isCast {
				return a // キャストできなくてpanicは避けたいのでそのまま返す
			}

			area.HotLevel = 1.0
			return area
		}).
		ToSlice(&calculated)

	return targetAreas, nil
}

func (model *HotLevelModelImpl) CalcSpotHotLevel(targetPins []types.SpotInfo) (calculated []types.SpotInfo, err error) {
	// TODO : いつか計算する
	From(targetPins).
		Select(func(a interface{}) interface{} {
			spot, isCast := a.(types.SpotInfo)
			if !isCast {
				return a // キャストできなくてpanicは避けたいのでそのまま返す
			}

			spot.HotLevel = 1.0
			return spot
		}).
		ToSlice(&calculated)

	return targetPins, nil
}
