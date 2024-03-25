package errno

const (
	Success             = 200
	InvalidParams       = 400
	Unauthorized        = 401
	InternalServerError = 500
)

// User
const (
	UserRegister       = 10000
	EncryptPassword    = 10001
	DatabaseCreateUser = 10002
	EmailFormat        = 10003

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
)
