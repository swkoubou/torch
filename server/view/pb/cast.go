package pb

import (
	"github.com/swkoubou/torch/server/model/types"
	"github.com/swkoubou/torch/server/view/pb/messages"
	"github.com/swkoubou/torch/server/view/pb/messages/structs"
)

// `cast.go` は、キャストに使うためだけの関数を置いておくための場所です

func ToSpotInfoResponse(spotInfo *types.SpotInfo) *messages.SpotInfoResponse {
	return &messages.SpotInfoResponse{
		SpotInfo: &structs.SpotInfo{
			Name:          spotInfo.Name,
			Description:   spotInfo.Description,
			PhotoFileName: spotInfo.PhotoFileName,
			Location: &structs.Location{
				Latitude:  spotInfo.Latitude,
				Longitude: spotInfo.Longitude,
			},
			EventSpan: &structs.TimeSpan{
				StartingAt: &structs.Time{
					Year:   int32(spotInfo.StartingAt.Year()),
					Month:  int32(spotInfo.StartingAt.Month()),
					Day:    int32(spotInfo.StartingAt.Day()),
					Hour:   int32(spotInfo.StartingAt.Hour()),
					Minute: int32(spotInfo.StartingAt.Minute()),
					Second: int32(spotInfo.StartingAt.Second()),
				},
				EndingAt: &structs.Time{
					Year:   int32(spotInfo.EndingAt.Year()),
					Month:  int32(spotInfo.EndingAt.Month()),
					Day:    int32(spotInfo.EndingAt.Day()),
					Hour:   int32(spotInfo.EndingAt.Hour()),
					Minute: int32(spotInfo.EndingAt.Minute()),
					Second: int32(spotInfo.EndingAt.Second()),
				},
			},
			SumLikeCount: uint32(spotInfo.Likes),
		},
	}
}
