package model

import (
	"github.com/swkoubou/torch/server/model/types"
	"mime/multipart"
)

// すべてのエリア情報を取得したり盛り上がり具合を計算するモデル
type AllAreasModel interface {
	// エリア情報をすべて取得
	Get() (areas []types.AreaInfo, err error)
	// エリア情報を盛り上がり具合を計算した上ですべて取得
	GetWithHotLevel() (areas []types.AreaInfo, err error)
}

// すべてのピン情報を取得したり盛り上がり具合を計算するモデル
type AllPinsModel interface {
	// エリア情報をすべて取得
	Get() (areas []types.PinInfo, err error)
	// エリア情報を盛り上がり具合を計算した上ですべて取得
	GetWithHotLevel() (areas []types.PinInfo, err error)
}

// エリアへのいいねと、いいね数の取得を行うモデル
type LikeAreaModel interface {
	Like(areaID uint) (count uint, err error)
	CountLikes(areaID uint) (count uint, err error)
}

// スポットへのいいねと、いいね数の取得を行うモデル
type LikeSpotModel interface {
	Like(spotID uint) (count uint, err error)
	CountLikes(spotID uint) (count uint, err error)
}

// 盛り上がり具合を計算するモデルです
type HotLevelModel interface {
	// エリアの盛り上がり具合を計算して、AreaInfoの配列に代入してから返す
	CalcAreaHotLevel(targetAreas []types.AreaInfo, targetPins []types.PinInfo) (calculated []types.AreaInfo, err error)
	// スポットの盛り上がり具合を計算して、PinInfoの配列に代入してから返す
	CalcPinHotLevel(targetPins []types.PinInfo) (calculated []types.PinInfo, err error)
}

// スポット情報を取得・追加するモデル
type SpotModel interface {
	// スポット情報をすべて取得
	Get(spotID uint) (spot *types.SpotInfo, err error)
	// スポット情報を追加
	Add(spot *types.NewSpotInfo) (err error)
	// スポットにいいねする
	Like(spotID uint) (err error)
}

// スポットの写真の保存・ファイル名を管理するモデル
type SpotPhotoModel interface {
	Add(image *multipart.FileHeader) (fileName string, err error) // ファイル名であってファイルパスではない
}
