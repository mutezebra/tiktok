package errno

var (
	SuccessErrno             = New(Success, "operate success")
	InvalidParamErrno        = New(InvalidParams, "invalid param")
	UnauthorizedErrno        = New(Unauthorized, "unauthorized")
	InternalServerErrorErrno = New(InternalServerError, "internal server failed")
)

// user
var (
	UserRegisterError           = New(UserRegister, "")
	EncryptPasswordError        = New(EncryptPassword, "encrypt password failed")
	DatabaseCreateUserError     = New(DatabaseCreateUser, "create user failed")
	EmailFormatError            = New(EmailFormat, "email format errno")
	DatabaseUserNameExistsError = New(DatabaseUserNameExists, "user name exists in database")

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

	TotpQrCodeError               = New(TotpQrCode, "")
	GenerateTotpError             = New(GenerateTotp, "generate totp qrcode failed")
	DatabaseUpdateTotpSecretError = New(DatabaseUpdateTotpSecret, "database update totp secret failed")

	EnableTotpError               = New(EnableTotp, "")
	DatabaseGetTotpSecretError    = New(DatabaseGetTotpSecret, "get totp secret from database failed")
	VerifyOtpCodeError            = New(VerifyOtpCode, "verify code with secret failed")
	DatabaseUpdateTotpStatusError = New(DatabaseUpdateTotpStatus, "update totp status in database failed")
)

// video
var (
	VideoFeedStreamSendError = New(VideoFeedStreamSend, "stream of video feed send resp failed")
	DatabaseGetVideoUrlError = New(DatabaseGetVideoUrl, "database get video url failed")
	OssGetVideoFeedError     = New(OssGetVideoFeed, "oss get video feed failed")
	IncrViewError            = New(IncrView, "incr view failed")

	OssUploadVideoError      = New(OssUploadVideo, "oss upload video failed")
	OssUploadVideoCoverError = New(OssUploadVideoCover, "oss upload video cover failed")
	DatabaseCreateVideoError = New(DatabaseCreateVideo, "database create video failed")
	OutOfLimitCoverSizeErrno = New(OutOfLimitCoverSize, "out of limit of cover size")

	DatabaseGetVideoListError = New(DatabaseGetVideoList, "database get video list failed")

	DatabaseSearchVideoError = New(DatabaseSearchVideo, "database search video failed")

	DatabaseGetVideoInfoError = New(DatabaseGetVideoInfo, "database get video info failed")
)

// interaction
var (
	DatabaseLikeCommentError                 = New(DatabaseLikeComment, "database like comment failed")
	DatabaseLikeVideoError                   = New(DatabaseLikeVideo, "database like video failed")
	DatabaseWhetherCommentLikeItemExistError = New(DatabaseWhetherCommentLikeItemExist, "database whether comment like item exist failed")
	CommentAlreadyLikedError                 = New(CommentAlreadyLiked, "comment already liked")
	DatabaseWhetherVideoLikeItemExistError   = New(DatabaseWhetherVideoLikeItemExist, "database if video like item exist failed")
	VideoAlreadyLikedError                   = New(VideoAlreadyLiked, "video already liked")
	DatabaseIfCommentExistError              = New(DatabaseIfCommentExist, "database if comment exist failed")
	CommentNotExistError                     = New(CommentNotExist, "comment not exist")
	DatabaseIfVideoExistError                = New(DatabaseIfVideoExist, "database if video exist failed")
	VideoNotExistError                       = New(VideoNotExist, "video not exist")

	DatabaseDislikeCommentError = New(DatabaseDislikeComment, "database dislike comment failed")
	DatabaseDislikeVideoError   = New(DatabaseDislikeVideo, "database dislike video failed")

	DatabaseLikeListError = New(DatabaseLikeList, "database get like list failed")

	DatabaseGetCommentRootIDError = New(DatabaseGetCommentRootID, "database get comment root id failed")
	DatabaseCreateCommentError    = New(DatabaseCreateComment, "database create comment failed")

	DatabaseGetCommentListError = New(DatabaseGetCommentList, "database get comment list failed")

	DatabaseDeleteCommentError = New(DatabaseDeleteComment, "database delete comment failed")
)

// chat
var (
	ChatRegisterError = New(ChatRegister, "chat register failed")
)

// relation
var (
	GroupNameTooLangError        = New(GroupNameTooLang, "group name too lang")
	GroupAlreadyExistError       = New(GroupAlreadyExist, "group already exist")
	DatabaseCreateChatGroupError = New(DatabaseCreateChatGroup, "database create chat group failed")

	UserNotExistError       = New(UserNotExist, "user not exist")
	DatabaseFollowError     = New(DatabaseFollow, "database follow failed")
	FollowAlreadyExistError = New(FollowAlreadyExist, "follow already exist")

	DatabaseGetFollowListError = New(DatabaseGetFollowList, "database get follow list failed")

	DatabaseGetFansListError = New(DatabaseGetFansList, "database get fans list failed")

	DatabaseGetFriendsListError = New(DatabaseGetFriendsList, "database get friends list failed")
)
