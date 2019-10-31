# pbRouter

WAFであるEchoで使えるProtocol Buffers Routerです。
@mox が作ったEchoでProtocol Buffersをいい感じに使うやつほとんどそのままです。
これもMVCのViewのヘルパーみたいな位置付けなので、
ソフトウェアの根幹っぽいディレクトリの外側に追い出すことにしました。


## 使い方

```go
e := echo.New()
e.Debug = true
	
// メッセージとハンドラはそれっぽいのを書いてね
message := proto.Message
handlerFunc := func(ctx echo.Context, pb proto.Message) (message proto.Message, err error){ return nil, nil}

apiGroup := NewProtocolBufferRouterByGroup(e.Group("/api"))
apiGroup.POST("/find", message, handlerFunc)

e.Logger.Fatal(e.Start(":1234"))
```

こういう感じで、Protocol Buffers特有のめんどくさいコードを
意識せずにルーターに追加することができるようになります。


## ToDo

(ToDo書く前にやったら？って感じですが、
このライブラリを完成させることは本質ではないので
このままにしておきます。)

* 認証が必要な通信をする関数で、ちゃんと認証できるようにする
    * 認証に使う関数をDIして使えるようにするとかで実現