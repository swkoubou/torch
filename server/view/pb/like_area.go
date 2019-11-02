package pb

import (
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
	"github.com/swkoubou/torch/server/model"
	"github.com/swkoubou/torch/server/view"
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
	panic("implement me")
}

func (view *LikeAreaPbViewImpl) GetPublishInterface() proto.Message {
	panic("implement me")
}
