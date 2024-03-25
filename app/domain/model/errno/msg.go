package errno

var (
	SuccessErrno             = New(Success, "operate success")
	InvalidParamErrno        = New(InvalidParams, "invalid param")
	UnauthorizedErrno        = New(Unauthorized, "unauthorized")
	InternalServerErrorErrno = New(InternalServerError, "internal server failed")
)

// user
var (
	UserRegisterError       = New(UserRegister, "")
	EncryptPasswordError    = New(EncryptPassword, "encrypt password failed")
	DatabaseCreateUserError = New(DatabaseCreateUser, "create user failed")
	EmailFormatError        = New(EmailFormat, "email format errno")

	UserLoginError               = New(UserLogin, "")
	GetPasswordFromDatabaseError = New(GetPasswordFromDatabase, "get password digest from db failed")
	CheckPasswordError           = New(CheckPassword, "checkout password failed")
	GenerateTokenError           = New(GenerateToken, "generate token failed")

	UserInfoError    = New(UserInfo, "")
	GetUserInfoError = New(GetUserInfo, "get user info failed")

	UserUploadAvatarError         = New(UserUploadAvatar, "")
	GetAvatarNameError            = New(GetAvatarName, "get avatar name failed")
	OssUploadAvatarError          = New(OssUploadAvatar, "oss upload avatar failed")
	DatabaseUpdateUserAvatarError = New(DatabaseUpdateUserAvatar, "database update user avatar failed")
	OutOfLimitAvatarSizeErrno     = New(OutOfLimitAvatarSize, "the avatar size out of the limit")

	DownloadAvatarError        = New(DownloadAvatar, "")
	DatabaseGetUserAvatarError = New(DatabaseGetUserAvatar, "get user avatar from database failed")
	OssDownloadAvatarError     = New(OssDownloadAvatar, "download avatar from oss failed")
)
