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

// すべてのスポット情報を取得したり盛り上がり具合を計算するモデル
type AllSpotsModel interface {
	// エリア情報をすべて取得
	Get() (spots []types.SpotInfo, err error)
	// エリア情報を盛り上がり具合を計算した上ですべて取得
	GetWithHotLevel() (spots []types.SpotInfo, err error)
}

// エリアへのいいねと、いいね数の取得を行うモデル
type LikeAreaModel interface {
	// いいねする
	Like(areaID uint) (err error)
	// いいねを数える
	CountLikes(areaID uint) (count uint, err error)
	// 指定されたすべてのエリアのいいねを数える
	// ぐるぐるSQLを避けるために関数を分けてある
	CountAllLikes(targetAreas []types.AreaInfo, targetSpots []types.SpotInfo) (counted []types.AreaInfo, err error)
}

// スポットへのいいねと、いいね数の取得を行うモデル
type LikeSpotModel interface {
	// いいねする
	Like(spotID uint) (err error)
	// いいねを数える
	CountLikes(spotID uint) (count uint, err error)
	// 指定されたすべてのスポットのいいねを数える
	// ぐるぐるSQLを避けるために関数を分けてある
	CountAllLikes(targetSpots []types.SpotInfo) (counted []types.SpotInfo, err error)
}

// 盛り上がり具合を計算するモデルです
type HotLevelModel interface {
	// エリアの盛り上がり具合を計算して、AreaInfoの配列に代入してから返す
	CalcAreaHotLevel(targetAreas []types.AreaInfo, targetSpots []types.SpotInfo) (calculated []types.AreaInfo, err error)
	// スポットの盛り上がり具合を計算して、SpotInfoの配列に代入してから返す
	CalcSpotHotLevel(targetSpots []types.SpotInfo) (calculated []types.SpotInfo, err error)
}

// スポット情報を取得・追加するモデル
type SpotModel interface {
	// スポット情報をすべて取得
	Get(spotID uint) (spot *types.SpotInfo, err error)
	// スポット情報を追加
	Add(spot *types.NewSpotInfo) (err error)
}

// スポットの写真の保存・ファイル名を管理するモデル
type SpotPhotoModel interface {
	// スポットの写真をファイル名が衝突しないように保存し、そのファイル名を返す
	// なお、「ファイル名」であってファイルパスではない
	Add(image *multipart.FileHeader) (fileName string, err error)
}
