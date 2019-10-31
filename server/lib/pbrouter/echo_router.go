package pbrouter

import (
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
)

type ProtocolBufferRouter struct {
	group *echo.Group
	echo  *echo.Echo
}

type PostHandlerFunc func(ctx echo.Context, pb proto.Message) (message proto.Message, err error)
type GetHandlerFunc func(ctx echo.Context) (message proto.Message, err error)

func NewProtocolBufferRouterByRouter(router *echo.Echo) *ProtocolBufferRouter {
	return &ProtocolBufferRouter{
		echo: router,
	}
}

func NewProtocolBufferRouterByGroup(group *echo.Group) *ProtocolBufferRouter {
	return &ProtocolBufferRouter{
		group: group,
	}
}

func (p *ProtocolBufferRouter) POST(relativePath string, value proto.Message, handler PostHandlerFunc) {
	p.registrationPost(relativePath, value, handler)
}

func (p *ProtocolBufferRouter) POSTWithAuth(relativePath string, value proto.Message, handler PostHandlerFunc) {
	p.registrationPost(relativePath, value, func(ctx echo.Context, pb proto.Message) (message proto.Message, err error) {
		token := ctx.Request().Header.Get("x-access-token")

		if token == "" {
			return nil, ctx.String(http.StatusUnauthorized, "no auth client")
		}

		//現状、問い合わせようがないのでコメントアウト
		//user, err := r.authRepo.GetUserWithDB(token)
		//if err != nil {
		//	ctx.String(http.StatusUnauthorized, "no auth client")
		//	return
		//}
		//ctx.Set("user", user)

		return nil, p.postCall(ctx, value, handler)
	})
}

func (p *ProtocolBufferRouter) GET(relativePath string, handler GetHandlerFunc) {
	p.registrationGet(relativePath, handler)
}

func (p *ProtocolBufferRouter) GETWithAuth(relativePath string, handler GetHandlerFunc) {
	p.registrationGet(relativePath, func(ctx echo.Context) (message proto.Message, err error) {
		token := ctx.Request().Header.Get("x-access-token")

		if token == "" {
			return nil, ctx.String(http.StatusUnauthorized, "no auth client")
		}

		//現状、問い合わせようがないのでコメントアウト
		//user, err := p.authRepo.GetUserWithDB(token)
		//if err != nil {
		//	return nil, ctx.String(http.StatusUnauthorized, "no auth client")
		//}
		//ctx.Set("user", user)

		return nil, p.getCall(ctx, handler)
	})
}

func (p *ProtocolBufferRouter) postCall(ctx echo.Context, value proto.Message, handler PostHandlerFunc) error {
	bodyBuff := ctx.Request().Body
	body, err := ioutil.ReadAll(bodyBuff)

	if err != nil {
		res, err := handler(ctx, nil)
		if err != nil {
			log.Printf("handler error: %+v\n", err)
			return ctx.String(http.StatusInternalServerError, "internal error")
		}

		resData, _ := proto.Marshal(res)
		return ctx.String(http.StatusOK, string(resData))
	}

	err = proto.Unmarshal(body, value)

	if err != nil {
		log.Printf("proto unmarshal error: %+v\n", err)
		res, err := handler(ctx, nil)
		if err != nil {
			log.Printf("handler error: %+v\n", err)
			return ctx.String(http.StatusInternalServerError, "internal error")
		}

		resData, _ := proto.Marshal(res)
		return ctx.String(http.StatusOK, string(resData))
	}

	res, err := handler(ctx, value)
	if err != nil {
		log.Printf("handler error: %+v\n", err)
		return ctx.String(http.StatusInternalServerError, "internal error")
	}

	resData, err := proto.Marshal(res)

	if err != nil {
		log.Printf("proto marshal error: %+v\n", err)
		return ctx.String(http.StatusInternalServerError, "internal error")
	}

	return ctx.String(http.StatusOK, string(resData))
}

func (p *ProtocolBufferRouter) getCall(ctx echo.Context, handler GetHandlerFunc) error {
	res, err := handler(ctx)
	if err != nil {
		log.Printf("handler error: %+v\n", err)
		return ctx.String(http.StatusInternalServerError, "internal error")
	}

	resData, err := proto.Marshal(res)

	if err != nil {
		log.Printf("proto marshal error: %+v\n", err)
		return ctx.String(http.StatusInternalServerError, "internal error")
	}

	return ctx.String(http.StatusOK, string(resData))
}

func (p *ProtocolBufferRouter) registrationGet(relativePath string, handler GetHandlerFunc) {
	if p.group != nil {
		p.group.GET(relativePath, func(ctx echo.Context) error {
			return p.getCall(ctx, handler)
		})
	} else if p.echo != nil {
		p.echo.GET(relativePath, func(ctx echo.Context) error {
			return p.getCall(ctx, handler)
		})
	} else {
		log.Fatal("UnknownRouterType")
	}
}

func (p *ProtocolBufferRouter) registrationPost(relativePath string, value proto.Message, handler PostHandlerFunc) {
	if p.group != nil {
		p.group.POST(relativePath, func(ctx echo.Context) error {
			return p.postCall(ctx, value, handler)
		})
	} else if p.echo != nil {
		p.echo.POST(relativePath, func(ctx echo.Context) error {
			return p.postCall(ctx, value, handler)
		})
	} else {
		log.Fatal("UnknownRouterType")
	}
}
