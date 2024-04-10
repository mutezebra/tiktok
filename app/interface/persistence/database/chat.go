package database

import (
	"context"
	"database/sql"
	"github.com/Mutezebra/tiktok/app/domain/repository"
	"github.com/Mutezebra/tiktok/pkg/log"
	"github.com/pkg/errors"
)

type ChatRepository struct {
	db *sql.DB
}

func NewChatRepository() *ChatRepository { return &ChatRepository{_db} }

func (repo *ChatRepository) WhetherExistUser(ctx context.Context, uid int64) (bool, error) {
	var exist bool
	err := repo.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM user WHERE id=?)", uid).Scan(&exist)
	return exist, err
}

func (repo *ChatRepository) CreateMessageWithChannel(ctx context.Context, msgs chan *repository.Message) {
	stmt, err := repo.db.PrepareContext(ctx, "INSERT INTO chat_messages(uid, receiver_id, content, create_at, delete_at) VALUES(?,?,?,?,?)")
	if err != nil {
		log.LogrusObj.Panic(errors.Wrap(err, "prepare query to create msg item failed"))
	}
	defer func() {
		if err = stmt.Close(); err != nil {
			log.LogrusObj.Error(errors.Wrap(err, "failed to close stmt"))
		}
	}()

	for {
		msg, ok := <-msgs
		if !ok {
			break
		}
		if _, err = stmt.ExecContext(ctx, msg.UID, msg.ReceiverID, msg.Content, msg.CreateAt, msg.DeleteAt); err != nil {
			log.LogrusObj.Error(errors.Wrap(err, "failed to create msg item"))
			break
		}
	}
}
