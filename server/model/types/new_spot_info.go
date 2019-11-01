package types

import "io"

// スポットを追加するときの構造体
type NewSpotInfo struct {
	Name        string
	Description string
	PhotoReader io.Reader
	LocatedAt   Location
}
