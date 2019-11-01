package pb

import (
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
	"github.com/swkoubou/torch/server/model"
	"github.com/swkoubou/torch/server/view"
)

type SpotInfoPbViewImpl struct {
	spotModel model.SpotModel
}

func NewSpotInfoPbView(spotModel model.SpotModel) (view view.SpotInfoPbView) {
	return &SpotInfoPbViewImpl{spotModel: spotModel}
}

func (view *SpotInfoPbViewImpl) GetPOSTHandler(ctx echo.Context, pb proto.Message) (message proto.Message, err error) {
	panic("implement me")
}

func (view *SpotInfoPbViewImpl) GetPublishInterface() proto.Message {
	panic("implement me")
}
