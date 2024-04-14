package usecase

import (
	"context"
	"github.com/Mutezebra/tiktok/app/domain/repository"

	"github.com/hertz-contrib/websocket"

	"github.com/Mutezebra/tiktok/app/domain/model/errno"
	"github.com/Mutezebra/tiktok/app/domain/service/chat"
	"github.com/Mutezebra/tiktok/app/interface/gateway/pack"
)

func ChatHandler(ctx context.Context, from, to int64, repo repository.ChatRepository) func(conn *websocket.Conn) {
	srv := chat.DefaultService(repo, true)
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
