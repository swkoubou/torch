package types

import (
	"mime/multipart"
	"time"
)

// スポットを追加するときの構造体
type NewSpotInfo struct {
	Name        string
	Description string
	PhotoReader *multipart.FileHeader
	LocatedAt   Location
	AreaID      uint

	StartingAt time.Time
	EndingAt   time.Time
}
