package usecase

import (
	"context"
	"github.com/Mutezebra/tiktok/app/domain/model/errno"
	"github.com/Mutezebra/tiktok/app/domain/service/chat"
	"github.com/Mutezebra/tiktok/app/interface/gateway/pack"
	"github.com/hertz-contrib/websocket"
)

func ChatHandler(ctx context.Context, from, to int64) func(conn *websocket.Conn) {
	srv := chat.DefaultService()
	return func(conn *websocket.Conn) {
		client := srv.NewClient(from, to, conn)
		if err := client.Register(ctx); err != nil {
			_ = pack.ReturnError(errno.ChatRegisterError, err)
			return
		}

		client.Read()
		client.Write()
	}
}
