package repository

import (
	"context"
)

// UserRepository defines the operational
// criteria for the user repository
type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error // create a new user
	UserNameExists(ctx context.Context, userName string) (bool, error)
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
	GetVideosInfo(ctx context.Context, vid []int64) ([]Video, error)
	GetVideoListByID(ctx context.Context, uid int64, page int, size int) ([]Video, error)
	SearchVideo(ctx context.Context, content string, page int, size int) ([]Video, error)
	GetVideoUrl(ctx context.Context, vid int64) (string, error)
	GetValByColumn(ctx context.Context, vid int64, column string) (string, error)
	UpdateViews(kvs map[int64]int32)
	GetVideoViews(ctx context.Context, vid int64) (int32, error)
}

type Video struct {
	ID       int64  `db:"id"`
	UID      int64  `db:"uid"`
	VideoURL string `db:"video_url"`
	CoverURL string `db:"cover_url"`
	Intro    string `db:"intro"`
	Title    string `db:"title"`
	VideoExt string `db:"video_ext"`
	CoverExt string `db:"cover_ext"`
	Starts   int32  `db:"starts"`
	Likes    int32  `db:"likes"`
	Views    int32  `db:"views"`
	CreateAt int64  `db:"create_at"`
	UpdateAt int64  `db:"update_at"`
	DeleteAt int64  `db:"delete_at"`
}

type InteractionRepository interface {
	CreateComment(ctx context.Context, comment *Comment) error
	LikeComment(ctx context.Context, uid, cid int64) error    // like a comment
	DislikeComment(ctx context.Context, uid, cid int64) error // dislike a comment
	DeleteComment(ctx context.Context, uid, cid int64) error  // delete a comment
	GetCommentRootID(ctx context.Context, cid int64) (int64, error)
	GetCommentList(ctx context.Context, cid int64, page int8, size int8) ([]Comment, error)
	WhetherCommentLikeItemExist(ctx context.Context, uid, cid int64) (bool, error) // check if the user has liked the video
	WhetherCommentExist(ctx context.Context, cid int64) (bool, error)

	LikeVideo(ctx context.Context, uid, vid int64) error                                               // like a video
	DislikeVideo(ctx context.Context, uid, vid int64) error                                            // dislike a video
	LikeList(ctx context.Context, uid int64, page int8, size int8) ([]Video, error)                    // get the user's like list
	GetVideoDirectCommentList(ctx context.Context, vid int64, page int8, size int8) ([]Comment, error) // get a video's comment list
	WhetherVideoLikeItemExist(ctx context.Context, uid, vid int64) (bool, error)                       // check if the user has liked the video
	WhetherVideoExist(ctx context.Context, vid int64) (bool, error)
}

type Comment struct {
	ID       int64  `db:"id"`
	UID      int64  `db:"uid"`
	VID      int64  `db:"vid"`
	RootID   int64  `db:"root_id"` // 用来标记这条评论是否是直接对视频的评论
	ReplyID  int64  `db:"reply_id"`
	Content  string `db:"content"`
	Likes    int32  `db:"likes"`
	CreateAt int64  `db:"create_at"`
	DeleteAt int64  `db:"delete_at"`
}

type ChatRepository interface {
	WhetherExistUser(ctx context.Context, uid int64) (bool, error)
	CreateMessage(ctx context.Context, msg *Message) error
	CreateMessageWithChannel(ctx context.Context, msgs chan *Message)
	ChatMessageHistory(ctx context.Context, req *HistoryQueryReq) ([]Message, error)
	NotReadMessage(ctx context.Context, uid int64, receiverID int64) ([]Message, error)
}

type Message struct {
	ID         int64  `db:"id" json:"id,omitempty"`
	UID        int64  `db:"uid" json:"uid,omitempty"`
	ReceiverID int64  `db:"receiver_id" json:"receiver_id,omitempty"`
	Content    string `db:"content" json:"content,omitempty"`
	HaveRead   bool   `db:"have_read" json:"have_read,omitempty"`
	CreateAt   int64  `db:"create_at" json:"create_at,omitempty"`
	DeleteAt   *int64 `db:"delete_at" json:"delete_at,omitempty"`
}

type HistoryQueryReq struct {
	PageSize   int8
	PageNum    int32
	Start, End int64
	From, To   int64
}

type RelationRepository interface {
	Follow(ctx context.Context, uid, followerID int64) error
	WhetherFollowExist(ctx context.Context, uid, followerID int64) (bool, error)
	GetFollowList(ctx context.Context, uid int64) ([]int64, error)
	GetFansList(ctx context.Context, uid int64) ([]int64, error)
	GetFriendList(ctx context.Context, uid int64) ([]int64, error)
	WhetherUserExist(ctx context.Context, uid int64) bool
}
