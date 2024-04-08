package errno

const (
	Success             = 200
	InvalidParams       = 400
	Unauthorized        = 401
	InternalServerError = 500
)

// User
const (
	UserRegister           = 10000
	EncryptPassword        = 10001
	DatabaseCreateUser     = 10002
	EmailFormat            = 10003
	DatabaseUserNameExists = 10004

	UserLogin               = 10010
	GetPasswordFromDatabase = 10011
	CheckPassword           = 10012
	GenerateToken           = 10013

	UserInfo    = 10020
	GetUserInfo = 10021

	UserUploadAvatar         = 10030
	GetAvatarName            = 10031
	VerifyAvatar             = 10032
	OssUploadAvatar          = 10033
	DatabaseUpdateUserAvatar = 10034
	OutOfLimitAvatarSize     = 10035

	DownloadAvatar        = 10040
	DatabaseGetUserAvatar = 10041
	OssDownloadAvatar     = 10042

	TotpQrCode               = 10050
	GenerateTotp             = 10051
	DatabaseUpdateTotpSecret = 10052

	EnableTotp               = 10060
	DatabaseGetTotpSecret    = 10061
	VerifyOtpCode            = 10062
	DatabaseUpdateTotpStatus = 10063
)

// Video
const (
	VideoFeedStreamSend = 20000
	DatabaseGetVideoUrl = 20001
	OssGetVideoFeed     = 20002
	IncrView            = 20003

	OssUploadVideo      = 20010
	OssUploadVideoCover = 20011
	DatabaseCreateVideo = 20012
	OutOfLimitCoverSize = 20013

	DatabaseGetVideoList = 20020

	DatabaseSearchVideo = 20030

	DatabaseGetVideoInfo = 20040
)

// Interaction
const (
	DatabaseLikeComment                 = 30000
	DatabaseLikeVideo                   = 30001
	DatabaseWhetherCommentLikeItemExist = 30002
	CommentAlreadyLiked                 = 30003
	DatabaseWhetherVideoLikeItemExist   = 30004
	VideoAlreadyLiked                   = 30005
	DatabaseIfCommentExist              = 30006
	CommentNotExist                     = 30007
	DatabaseIfVideoExist                = 30008
	VideoNotExist                       = 30009

	DatabaseDislikeComment = 30010
	DatabaseDislikeVideo   = 30011

	DatabaseLikeList = 30020

	DatabaseGetCommentRootID = 30030
	DatabaseCreateComment    = 30031

	DatabaseGetCommentList = 30040

	DatabaseDeleteComment = 30050
)

// Chat
const (
	ChatRegister = 40000
)
