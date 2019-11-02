package main

import (
	"github.com/jinzhu/gorm"
	"github.com/swkoubou/torch/server/model/types"
)

// プリセット的なデータの挿入
func Prepare(db *gorm.DB) {
	// エリア情報がカラッポだったらプリセットのデータを挿入
	var numAllergies uint = 0
	db.Table("area_infos").Count(&numAllergies)
	if numAllergies != 0 {
		return
	}

	// エリア情報の追加
	db.Create(&types.AreaInfo{Name: "自動車エリア", LeftUpX: 35.487920, LeftUpY: 139.340501, RightBottomX: 35.487358, RightBottomY: 139.341577})
	db.Create(&types.AreaInfo{Name: "応用バイオエリア", LeftUpX: 35.487358, LeftUpY: 139.340501, RightBottomX: 35.486804, RightBottomY: 139.341387})
	db.Create(&types.AreaInfo{Name: "図書館エリア", LeftUpX: 35.487357, LeftUpY: 139.341387, RightBottomX: 35.486804, RightBottomY: 139.342138})
	db.Create(&types.AreaInfo{Name: "機械工学エリア", LeftUpX: 35.487359, LeftUpY: 139.342139, RightBottomX: 35.486997, RightBottomY: 139.342721})
	db.Create(&types.AreaInfo{Name: "工学エリア", LeftUpX: 35.488034, LeftUpY: 139.341578, RightBottomX: 35.487358, RightBottomY: 139.342139})
	db.Create(&types.AreaInfo{Name: "ロボットエリア", LeftUpX: 35.487700, LeftUpY: 139.342139, RightBottomX: 35.487201, RightBottomY: 139.343700})
	db.Create(&types.AreaInfo{Name: "学生サービス棟エリア", LeftUpX: 35.486737, LeftUpY: 139.340841, RightBottomX: 35.486324, RightBottomY: 139.341524})
	db.Create(&types.AreaInfo{Name: "情報学エリア", LeftUpX: 35.486324, LeftUpY: 139.340899, RightBottomX: 35.485898, RightBottomY: 139.341524})
	db.Create(&types.AreaInfo{Name: "講義棟エリア", LeftUpX: 35.485898, LeftUpY: 139.340899, RightBottomX: 35.485463, RightBottomY: 139.341743})
	db.Create(&types.AreaInfo{Name: "KAITHALLエリア", LeftUpX: 35.485463, LeftUpY: 139.340897, RightBottomX: 35.485247, RightBottomY: 139.341743})
}