package messages

import (
	"github.com/swkoubou/torch/server/model/types"
	"time"
)

type NewSpotInfoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	LocLatitude  float64 `json:"loc_latitude"`
	LocLongitude float64 `json:"loc_longitude"`

	AreaID uint `json:"area_id"`

	StartYear   int `json:"start_year"`
	StartMonth  int `json:"start_month"`
	StartDay    int `json:"start_day"`
	StartHour   int `json:"start_hour"`
	StartMinute int `json:"start_minute"`
	StartSecond int `json:"start_second"`

	EndYear   int `json:"end_year"`
	EndMonth  int `json:"end_month"`
	EndDay    int `json:"end_day"`
	EndHour   int `json:"end_hour"`
	EndMinute int `json:"end_minute"`
	EndSecond int `json:"end_second"`
}

func (req *NewSpotInfoRequest) GetLocation() *types.Location {
	return &types.Location{
		Latitude:  req.LocLatitude,
		Longitude: req.LocLongitude,
	}
}

func (req *NewSpotInfoRequest) GetStartTime() time.Time {
	return time.Date(
		req.StartYear,
		time.Month(req.StartMonth),
		req.StartDay,
		req.StartHour,
		req.StartMinute,
		req.StartSecond,
		0,
		time.Local,
	)
}

func (req *NewSpotInfoRequest) GetEndTime() time.Time {
	return time.Date(
		req.EndYear,
		time.Month(req.EndMonth),
		req.EndDay,
		req.EndHour,
		req.EndMinute,
		req.EndSecond,
		0,
		time.Local,
	)
}
