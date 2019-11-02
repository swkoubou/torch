package pb

import (
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
	"github.com/swkoubou/torch/server/model"
	"github.com/swkoubou/torch/server/view"
	"github.com/swkoubou/torch/server/view/pb/messages"
)

type LikeSpotPbViewImpl struct {
	likeSpotModel model.LikeSpotModel
}

func NewLikeSpotPbView(likeSpotModel model.LikeSpotModel) view.LikeSpotPbView {
	return &LikeSpotPbViewImpl{
		likeSpotModel: likeSpotModel,
	}
}

func (view *LikeSpotPbViewImpl) GetPOSTHandler(ctx echo.Context, pb proto.Message) (message proto.Message, err error) {
	// リクエストの解釈
	requestPb := pb.(*messages.SpotLikeRequest)
	reqSpotId := uint(requestPb.SpotId)

	// モデルに投げる
	err = view.likeSpotModel.Like(reqSpotId)
	if err != nil {
		return nil, err
	}

	// レスポンスをつくる
	response := &messages.SpotLikeResponse{Message: "success"}
	return response, nil
}

func (view *LikeSpotPbViewImpl) GetPublishInterface() proto.Message {
	return &messages.SpotLikeRequest{}
}
