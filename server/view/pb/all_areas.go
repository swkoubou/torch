package pb

import (
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
	"github.com/swkoubou/torch/server/model"
	"github.com/swkoubou/torch/server/view"
)

type AllAreasPbViewImpl struct {
	allAreasModel model.AllAreasModel
}

func NewAllAreasPbView(allAreasModel model.AllAreasModel) view.AllAreasPbView {
	return &AllAreasPbViewImpl{allAreasModel: allAreasModel}
}

func (view *AllAreasPbViewImpl) GetGETHandler(ctx echo.Context) (message proto.Message, err error) {
	// モデルに投げる
	founds, err := view.allAreasModel.GetWithHotLevel()
	if err != nil {
		return nil, err
	}

	// レスポンスをつくる
	response := ToAllAreaInfoResponse(founds)
	return response, nil
}
