package router

import (
	"context"
	"github.com/Mutezebra/tiktok/app/interface/gateway/middleware"

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
		user.POST("/register", handler.UserRegisterHandler())
		user.POST("/login", handler.UserLoginHandler())

		auth := user.Group("/auth")
		auth.Use(middleware.JWT())
		{
			auth.GET("/info", handler.UserInfoHandler())
			auth.POST("/upload-avatar", handler.UploadAvatarHandler())
			auth.GET("/download-avatar", handler.DownloadAvatarHandler())
			auth.GET("/totp-qrcode", handler.TotpQRCodeHandler())
			auth.POST("/enable-totp", handler.EnableTotpHandler())
		}
	}

	video := v1.Group("/video")
	{
		video.GET("/feed", handler.VideoFeedHandler())
		video.GET("/popular", handler.VideoPopularHandler())
		video.POST("/search", handler.VideoSearchHandler())

		auth := video.Group("/auth")
		auth.Use(middleware.JWT())
		{
			auth.POST("/publish", handler.VideoPublishHandler())
			auth.GET("/list", handler.VideoListHandler())
		}
	}

	return h
}
