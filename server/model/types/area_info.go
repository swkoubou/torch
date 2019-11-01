package types

import "github.com/jinzhu/gorm"

type AreaInfo struct {
	gorm.Model

	Name  string
	Likes uint

	// Region
	// DBで変にFK作られたり構造体に入れ直そうとするとエグいので切り出す
	LeftUpX      float64
	LeftUpY      float64
	RightBottomX float64
	RightBottomY float64

	HotLevel float64 `gorm:"-"`
}

func (area *AreaInfo) GetRegion() Rectangle {
	rect := Rectangle{
		LeftUp: Location{
			Latitude:  area.LeftUpX,
			Longitude: area.LeftUpY,
		},
		RightBottom: Location{
			Latitude:  area.RightBottomX,
			Longitude: area.RightBottomY,
		},
	}

	return rect
}
