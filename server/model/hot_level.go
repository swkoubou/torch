package model

import (
	"github.com/swkoubou/torch/server/model/types"
)

type HotLevelModelImpl struct{}

func NewHotLevelModel() HotLevelModel {
	return &HotLevelModelImpl{}
}

func (model *HotLevelModelImpl) CalcAreaHotLevel(targetAreas []types.AreaInfo, targetPins []types.PinInfo) (calculated []types.AreaInfo, err error) {
	// TODO : いつか計算する
	return targetAreas, nil
}

func (model *HotLevelModelImpl) CalcPinHotLevel(targetPins []types.PinInfo) (calculated []types.PinInfo, err error) {
	// TODO : いつか計算する
	return targetPins, nil
}
