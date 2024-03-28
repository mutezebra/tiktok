package repository

import (
	"context"
)

// UserRepository defines the operational
// criteria for the user repository
type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetPasswordAndIDByName(ctx context.Context, name string) (string, int64, error)
	UserInfoByID(ctx context.Context, id int64) (*User, error)
	UserInfoByName(ctx context.Context, name string) (*User, error)
	UpdateUserAvatar(ctx context.Context, filename string, uid int64) error
	GetUserAvatar(ctx context.Context, uid int64) (string, error)
	UpdateTotpSecret(ctx context.Context, uid int64, secret string) error
	GetTotpSecret(ctx context.Context, uid int64) (string, error)
	UpdateTotpStatus(ctx context.Context, status bool, uid int64) error
	UpdateColumnByKV(ctx context.Context, uid int64, k, v string) error
	GetColumnByKUID(ctx context.Context, key string, uid int64) (string, error)
}

// User is the standards for repo operand objects
type User struct {
	ID             int64  `db:"id"`
	UserName       string `db:"user_name"`
	Email          string `db:"email"`
	PasswordDigest string `db:"password_digest"`
	Gender         int8   `db:"gender"`
	Avatar         string `db:"avatar"`
	Fans           int32  `db:"fans"`
	Follows        int32  `db:"follows"`
	TotpEnable     bool   `db:"totp_enable"`
	TotpSecret     string `db:"totp_secret"`
	CreateAt       int64  `db:"create_at"`
	UpdateAt       int64  `db:"update_at"`
	DeleteAt       int64  `db:"delete_at"`
}

type VideoRepository interface {
	CreateVideo(ctx context.Context, video *Video) (int64, error)
	GetVideoInfo(ctx context.Context, vid int64) (*Video, error)
	GetVideoListByID(ctx context.Context, uid int64, page int, size int) ([]Video, error)
	GetVideoPopular(ctx context.Context, vid []int64) ([]Video, error)
	SearchVideo(ctx context.Context, content string, page int, size int) ([]Video, error)
	GetVideoUrl(ctx context.Context, vid int64) (string, error)
	GetValByColumn(ctx context.Context, vid int64, column string) (string, error)
}

type Video struct {
	ID        int64  `db:"id"`
	UID       int64  `db:"uid"`
	VideoURL  string `db:"video_url"`
	CoverURL  string `db:"cover_url"`
	Intro     string `db:"intro"`
	Title     string `db:"title"`
	VideoExt  string `db:"video_ext"`
	CoverExt  string `db:"cover_ext"`
	Starts    int32  `db:"starts"`
	Favorites int32  `db:"favorites"`
	Views     int32  `db:"views"`
	CreateAt  int64  `db:"create_at"`
	UpdateAt  int64  `db:"update_at"`
	DeleteAt  int64  `db:"delete_at"`
}
