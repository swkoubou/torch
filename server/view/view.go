package view

import (
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
)

// エリアをすべて取得する
type AllAreasPbView interface {
	// GETのハンドラー
	GetGETHandler(ctx echo.Context) (message proto.Message, err error)
}

// スポットをすべて取得する
type AllSpotsPbView interface {
	// GETのハンドラー
	GetGETHandler(ctx echo.Context) (message proto.Message, err error)
}

// エリアへのいいねを行う
type LikeAreaPbView interface {
	// POSTのハンドラー
	GetPOSTHandler(ctx echo.Context, pb proto.Message) (message proto.Message, err error)
	// POSTで受け取る proto.Message を埋め込んだ構造体を返す
	GetPublishInterface() proto.Message
}

// スポットへのいいねを行う
type LikeSpotPbViewModel interface {
	// POSTのハンドラー
	GetPOSTHandler(ctx echo.Context, pb proto.Message) (message proto.Message, err error)
	// POSTで受け取る proto.Message を埋め込んだ構造体を返す
	GetPublishInterface() proto.Message
}

// スポットの取得を行う
type SpotInfoPbView interface {
	// POSTのハンドラー
	GetPOSTHandler(ctx echo.Context, pb proto.Message) (message proto.Message, err error)
	// POSTで受け取る proto.Message を埋め込んだ構造体を返す
	GetPublishInterface() proto.Message
}

// (管理用)スポットの追加を行う `multipart/form-data` なハンドラー
type SpotInfoRESTView interface {
	// POSTのハンドラー
	GetPUTHandler(ctx echo.Context) (err error)
}
