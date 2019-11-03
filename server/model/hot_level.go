package model

import (
	"github.com/swkoubou/torch/server/model/types"
	"gonum.org/v1/gonum/stat"
)

type HotLevelModelImpl struct{}

func NewHotLevelModel() HotLevelModel {
	return &HotLevelModelImpl{}
}

func (model *HotLevelModelImpl) CalcAreaHotLevel(targetAreas []types.AreaInfo, targetSpots []types.SpotInfo) (calculated []types.AreaInfo, err error) {
	var likeArray []float64
	for _, area := range targetAreas {
		likeArray = append(likeArray, float64(area.Likes))
	}

	mean, variance := stat.MeanVariance(likeArray, nil)

	for _, area := range targetAreas {
		area.HotLevel = model.calcHotLevel(float64(area.Likes), mean, variance, 1)
		calculated = append(calculated, area)
	}

	return calculated, nil
}

func (model *HotLevelModelImpl) CalcSpotHotLevel(targetSpots []types.SpotInfo) (calculated []types.SpotInfo, err error) {
	var data []float64
	for _, spot := range targetSpots {
		data = append(data, float64(spot.Likes))
	}

	mean, variance := stat.MeanVariance(data, nil)

	for _, spot := range targetSpots {
		gravity := model.calcSpotGravity(spot)

		spot.HotLevel = model.calcHotLevel(float64(spot.Likes), mean, variance, gravity)
		calculated = append(calculated, spot)
	}

	return calculated, nil
}

func (model *HotLevelModelImpl) calcSpotGravity(info types.SpotInfo) float64 {
	duration := info.GetHourSpan()
	rawGrav := -0.2173913*duration + 6.2173913

	return model.alignMinMax(rawGrav, 1, 6)
}

func (model *HotLevelModelImpl) alignMinMax(target, min, max float64) float64 {
	if min > target {
		return min
	}
	if max < target {
		return max
	}

	return target
}

func (model *HotLevelModelImpl) calcHotLevel(likes, mean, variance, gravity float64) float64 {
	rawLevel := (likes - (mean - 1.5*variance)) / (3 * variance) * 100
	level := gravity * rawLevel

	return model.alignMinMax(level, 0, 100)
}
