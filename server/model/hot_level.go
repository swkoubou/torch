package model

import (
	"github.com/swkoubou/torch/server/model/types"
)

type HotLevelModelImpl struct{}

func NewHotLevelModel() HotLevelModel {
	return &HotLevelModelImpl{}
}

func (model *HotLevelModelImpl) CalcAreaHotLevel(targetAreas []types.AreaInfo, targetSpots []types.SpotInfo) (calculated []types.AreaInfo, err error) {
	for _, area := range targetAreas {
		area.HotLevel = 1.0
		calculated = append(calculated, area)
	}

	return targetAreas, nil
}

func (model *HotLevelModelImpl) CalcSpotHotLevel(targetSpots []types.SpotInfo) (calculated []types.SpotInfo, err error) {
	for _, spot := range targetSpots {
		spot.HotLevel = 1.0
		calculated = append(calculated, spot)
	}

	return targetSpots, nil
}
