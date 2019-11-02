package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/swkoubou/torch/server/lib/config"
)

func main() {
	// サーバーのインスタンスの準備
	e := echo.New()

	// 設定の読み込み
	config, err := config.Load()
	if err != nil {
		e.Logger.Fatal("Can't load config.", err)
		return
	}

	// DBのインスタンスの準備
	dbLoc := config.GetDBServerLocation()
	db, err := ConnectToDB(dbLoc)
	if err != nil {
		e.Logger.Fatal("Can't connect to db.", err)
		return
	}

	// オートマイグレーション
	err = Migrate(db)
	if err != nil {
		e.Logger.Fatal("Can't migrate models to db.", err)
		return
	}

	// エリア情報の追加
	Prepare(db)

	// ルーティング
	err = Route(e, db)
	if err != nil {
		e.Logger.Fatal("Can't set handlers to router.", err)
		return
	}

	// サーバーの起動
	apiLoc := config.GetAPIServerLocation()
	e.Logger.Fatal(e.Start(apiLoc))
}
