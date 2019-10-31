package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

const ApiServerStartingPort = "TORCH_API_PORT"

func main() {
	// ポートの読み込み
	port := os.Getenv(ApiServerStartingPort)
	if port == "" {
		log.Fatalln("Can't load starting port.")
		return
	}

	// サーバーのインスタンスの準備
	e := echo.New()

	// ルーティング
	// これはテスト用にHello, Worldしてる
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// サーバーの起動
	location := ":" + port
	e.Logger.Fatal(e.Start(location))
}
