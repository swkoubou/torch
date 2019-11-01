package pb

import (
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
	"github.com/swkoubou/torch/server/model"
	"github.com/swkoubou/torch/server/view"
	"github.com/swkoubou/torch/server/view/pb/messages"
)

type SpotInfoPbViewImpl struct {
	spotModel model.SpotModel
}

func NewSpotInfoPbView(spotModel model.SpotModel) (view view.SpotInfoPbView) {
	return &SpotInfoPbViewImpl{spotModel: spotModel}
}

func (view *SpotInfoPbViewImpl) GetPOSTHandler(ctx echo.Context, pb proto.Message) (message proto.Message, err error) {
	// リクエストの解釈
	requestPb := pb.(*messages.SpotInfoRequest)
	reqSpotId := uint(requestPb.GetSpotId())

	// モデルに投げる
	found, err := view.spotModel.Get(reqSpotId)
	if err != nil {
		return nil, err
	}

	// レスポンスをつくる
	response := ToSpotInfoResponse(found)
	return response, nil
}

func (view *SpotInfoPbViewImpl) GetPublishInterface() proto.Message {
	return &messages.SpotInfoRequest{}
}
