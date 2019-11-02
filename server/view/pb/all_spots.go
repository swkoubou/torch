package pb

import (
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
	"github.com/swkoubou/torch/server/model"
	"github.com/swkoubou/torch/server/view"
)

type AllSpotsPbViewImpl struct {
	allSpotsModel model.AllSpotsModel
}

func NewAllSpotsPbView(allSpotsModel model.AllSpotsModel) view.AllSpotsPbView {
	return &AllSpotsPbViewImpl{allSpotsModel: allSpotsModel}
}

func (view *AllSpotsPbViewImpl) GetGETHandler(ctx echo.Context) (message proto.Message, err error) {
	// モデルに投げる
	founds, err := view.allSpotsModel.GetWithHotLevel()
	if err != nil {
		return nil, err
	}

	// レスポンスをつくる
	response := ToAllSpotInfoResponse(founds)
	return response, nil
}
