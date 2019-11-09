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
	// いいねだけを取り出した配列(平均と分散の計算に使う)
	var likesArray []float64
	for _, area := range targetAreas {
		likesArray = append(likesArray, float64(area.Likes))
	}

	// いいねの標本平均、標本分散をとる
	mean, variance := stat.MeanVariance(likesArray, nil)

	// 各エリアの盛り上がり度合いを計算し、詰め直して返す
	for _, area := range targetAreas {
		area.HotLevel = model.calcHotLevel(float64(area.Likes), mean, variance, 3)
		calculated = append(calculated, area)
	}

	return calculated, nil
}

func (model *HotLevelModelImpl) CalcSpotHotLevel(targetSpots []types.SpotInfo) (calculated []types.SpotInfo, err error) {
	// いいねだけを取り出した配列(平均と分散の計算に使う)
	var likesArray []float64
	for _, spot := range targetSpots {
		likesArray = append(likesArray, float64(spot.Likes))
	}

	// いいねの標本平均、標本分散をとる
	mean, variance := stat.MeanVariance(likesArray, nil)

	// 各スポットの盛り上がり度合いを計算し、詰め直して返す
	for _, spot := range targetSpots {
		gravity := model.calcSpotGravity(spot)

		spot.HotLevel = model.calcHotLevel(float64(spot.Likes), mean, variance, gravity)
		calculated = append(calculated, spot)
	}

	return calculated, nil
}

// スポットの盛り上がり度合いを、イベントの開催期間の短さで重みをつける
func (model *HotLevelModelImpl) calcSpotGravity(info types.SpotInfo) float64 {
	duration := info.GetHourSpan()
	rawGrav := -0.2173913*duration + 6.2173913

	return model.alignMinMax(rawGrav, 1, 3)
}

// データを最大値・最小値を指定して丸める
func (model *HotLevelModelImpl) alignMinMax(target, min, max float64) float64 {
	if min > target {
		return min
	}
	if max < target {
		return max
	}

	return target
}

// 盛り上がり度合いを算出する
func (model *HotLevelModelImpl) calcHotLevel(likes, mean, variance, gravity float64) float64 {
	// とりあえず、いいね数の偏差値の中央値を10とし、振れ幅を50倍にした値
	// TODO : 偏差値だけだと盛り上がりがいい感じに計測できないので、なんとかする
	k := 1.0
	rawLevel := 500*(likes-mean)/(k*variance) + 10
	level := gravity * rawLevel

	return model.alignMinMax(level, 0, 100)
}
