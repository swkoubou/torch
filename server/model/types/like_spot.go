package types

import "github.com/jinzhu/gorm"

type LikeSpot struct {
	gorm.Model

	SpotInfoID uint
	SpotInfo   SpotInfo
}
