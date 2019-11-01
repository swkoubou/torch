package types

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SpotInfo struct {
	gorm.Model

	Name          string
	Description   string
	PhotoFileName string
	Likes         uint

	// Location
	// DBで変にFK作られたり構造体に入れ直そうとするとエグいので切り出す
	Latitude  float64
	Longitude float64

	AreaInfoID uint
	AreaInfo   AreaInfo

	StartingAt time.Time
	EndingAt   time.Time

	HotLevel float64 `gorm:"-"`
}

func (spot *SpotInfo) GetLocation() Location {
	return Location{
		Latitude:  spot.Latitude,
		Longitude: spot.Longitude,
	}
}
