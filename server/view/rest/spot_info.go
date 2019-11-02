package rest

import (
	"github.com/labstack/echo"
	"github.com/swkoubou/torch/server/model"
	"github.com/swkoubou/torch/server/view"
)

type SpotInfoRESTViewImpl struct {
	spotModel model.SpotModel
}

func NewSpotInfoRESTView(spotModel model.SpotModel) view.SpotInfoRESTView {
	return &SpotInfoRESTViewImpl{
		spotModel: spotModel,
	}
}

func (view *SpotInfoRESTViewImpl) GetPUTHandler(ctx echo.Context) (err error) {
	panic("implement me")
}
