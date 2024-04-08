package database

import (
	"context"
	"database/sql"
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
