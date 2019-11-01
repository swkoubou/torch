package types

import "github.com/jinzhu/gorm"

type LikeArea struct {
	gorm.Model

	AreaInfoID uint
	AreaInfo   AreaInfo
}
