package types

import "time"

// ピンの情報
// 内部的にはScanで実装する ( http://gorm.io/docs/query.html#Scan )
type PinInfo struct {
	ID    uint
	Likes uint `gorm:"-"`

	// Location
	// DBで変にFK作られたり構造体に入れ直そうとするとエグいので切り出す
	Latitude  float64
	Longitude float64

	AreaID   uint
	AreaInfo AreaInfo

	StartingAt time.Time
	EndingAt   time.Time

	HotLevel float64 `gorm:"-"`
}

func (pin *PinInfo) GetLocation() Location {
	return Location{
		Latitude:  pin.Latitude,
		Longitude: pin.Longitude,
	}
}
