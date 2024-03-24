package router

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/Mutezebra/tiktok/app/interface/gateway/handler"
	"github.com/Mutezebra/tiktok/config"
	"github.com/Mutezebra/tiktok/consts"
)

func NewRouter() *server.Hertz {
	h := server.Default(
		server.WithHostPorts(config.Conf.Service[consts.GatewayServiceKey].Address),
		server.WithMaxRequestBodySize(100*consts.MB),
	)
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.String(200, "pong")
	})
	v1 := h.Group("/api")
	user := v1.Group("/user")
	{
		user.GET("/register", handler.UserRegisterHandler())
		user.GET("/login", handler.UserLoginHandler())
	}

	return h
}
