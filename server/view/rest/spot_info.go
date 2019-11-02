package rest

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/swkoubou/torch/server/model"
	"github.com/swkoubou/torch/server/model/types"
	"github.com/swkoubou/torch/server/view"
	"github.com/swkoubou/torch/server/view/rest/messages"
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
	// リクエストを読む(画像)
	fileInfo, err := ctx.FormFile("image")
	if err != nil {
		_ = ctx.String(401, "failure")
		fmt.Printf("SpotInfoRESTView.PUT(): %v\n", err)
		return
	}

	// リクエストを読む(JSON)
	optionJson := ctx.FormValue("option")
	if err != nil {
		_ = ctx.String(401, "failure")
		fmt.Printf("SpotInfoRESTView.PUT(): %v\n", err)
		return
	}

	/// JSONのUnmarshal
	optionStruct := &messages.NewSpotInfoRequest{}
	err = json.Unmarshal([]byte(optionJson), optionStruct)
	if err != nil {
		_ = ctx.String(401, "failure")
		fmt.Printf("SpotInfoRESTView.PUT(): %v\n", err)
		return
	}

	// 内部の型に書き換え
	newSpotInfo := &types.NewSpotInfo{
		Name:        optionStruct.Name,
		Description: optionStruct.Description,
		PhotoReader: fileInfo,
		LocatedAt:   *optionStruct.GetLocation(),
		AreaID:      optionStruct.AreaID,
		StartingAt:  optionStruct.GetStartTime(),
		EndingAt:    optionStruct.GetEndTime(),
	}

	err = view.spotModel.Add(newSpotInfo)
	if err != nil {
		_ = ctx.String(401, "failure")
		fmt.Printf("SpotInfoRESTView.PUT(): %v\n", err)
		return
	}

	_ = ctx.String(200, "success")
	return nil
}
