package database

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Mutezebra/tiktok/app/domain/repository"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{_db}
}

// CreateUser create a repository.User object in database.
func (repo *UserRepository) CreateUser(ctx context.Context, user *repository.User) error {
	_, err := repo.db.ExecContext(ctx,
		"INSERT INTO user(id,user_name,email,password_digest,gender,avatar,fans,follows,totp_enable,totp_secret,create_at,update_at,delete_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)",
		user.ID, user.UserName, user.Email, user.PasswordDigest,
		user.Gender, user.Avatar, user.Fans, user.Follows,
		user.TotpEnable, user.TotpSecret, user.CreateAt, user.UpdateAt,
		user.DeleteAt)

	return err
}

func (repo *UserRepository) GetPasswordAndIDByName(ctx context.Context, name string) (string, int64, error) {
	var passwordDigest string
	var id int64
	if err := repo.db.QueryRowContext(ctx, "SELECT password_digest,id from user WHERE user_name=? LIMIT 1", name).Scan(
		&passwordDigest, &id); err != nil {
		return "", 0, err
	}

	return passwordDigest, id, nil
}

// UserInfoByID retrieves a user's information from the database using the user's ID.
func (repo *UserRepository) UserInfoByID(ctx context.Context, id int64) (*repository.User, error) {
	var user repository.User
	if err := repo.db.QueryRowContext(ctx, "SELECT * from user WHERE id=? LIMIT 1", id).Scan(
		&user.ID, &user.UserName, &user.Email, &user.PasswordDigest,
		&user.Gender, &user.Avatar, &user.Fans, &user.Follows,
		&user.TotpEnable, &user.TotpSecret, &user.CreateAt, &user.UpdateAt,
		&user.DeleteAt); err != nil {
		return nil, err
	}

	return &user, nil
}

// UserInfoByName retrieves a user's information from the database using the user's ID.
func (repo *UserRepository) UserInfoByName(ctx context.Context, name string) (*repository.User, error) {
	var user repository.User
	if err := repo.db.QueryRowContext(ctx, "SELECT * from user WHERE user_name=? LIMIT 1", name).Scan(
		&user.ID, &user.UserName, &user.Email, &user.PasswordDigest,
		&user.Gender, &user.Avatar, &user.Fans, &user.Follows,
		&user.TotpEnable, &user.TotpSecret, &user.CreateAt, &user.UpdateAt,
		&user.DeleteAt); err != nil {
		return nil, err
	}

	return &user, nil
}
