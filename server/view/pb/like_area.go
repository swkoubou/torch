package pb

import (
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
	"github.com/swkoubou/torch/server/model"
	"github.com/swkoubou/torch/server/view"
	"github.com/swkoubou/torch/server/view/pb/messages"
)

type LikeAreaPbViewImpl struct {
	likeAreaModel model.LikeAreaModel
}

func NewLikeAreaPbView(likeAreaModel model.LikeAreaModel) view.LikeAreaPbView {
	return &LikeAreaPbViewImpl{
		likeAreaModel: likeAreaModel,
	}
}

func (view *LikeAreaPbViewImpl) GetPOSTHandler(ctx echo.Context, pb proto.Message) (message proto.Message, err error) {
	// リクエストの解釈
	requestPb := pb.(*messages.AreaLikeRequest)
	reqAreaId := uint(requestPb.AreaId)

	// モデルに投げる
	err = view.likeAreaModel.Like(reqAreaId)
	if err != nil {
		return nil, err
	}

	// レスポンスをつくる
	response := &messages.AreaLikeResponse{Message: "success"}
	return response, nil
}

func (view *LikeAreaPbViewImpl) GetPublishInterface() proto.Message {
	return &messages.AreaLikeRequest{}
}
