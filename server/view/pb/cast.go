package pb

import (
	. "github.com/ahmetb/go-linq"
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

func ToAllAreaInfoResponse(areaInfos []types.AreaInfo) *messages.AllAreaInfoResponse {
	var pbAreaInfos []*structs.AreaInfo

	From(areaInfos).
		Select(func(a interface{}) interface{} {
			area := a.(types.AreaInfo)
			return ToPbAreaInfo(&area)
		}).
		ToSlice(pbAreaInfos)

	return &messages.AllAreaInfoResponse{
		AreaInfos: pbAreaInfos,
	}
}

func ToPbAreaInfo(areaInfos *types.AreaInfo) *structs.AreaInfo {
	return &structs.AreaInfo{
		AreaId: uint32(areaInfos.Model.ID),
		Name:   areaInfos.Name,
		Region: &structs.Rectangle{
			LeftUp: &structs.Location{
				Latitude:  areaInfos.LeftUpX,
				Longitude: areaInfos.LeftUpY,
			},
			RightBottom: &structs.Location{
				Latitude:  areaInfos.RightBottomX,
				Longitude: areaInfos.RightBottomY,
			},
		},
		HotScore:          areaInfos.HotLevel,
		SpecificLikeCount: uint32(areaInfos.Likes),
	}
}
