package errno

var (
	SuccessErrno             = New(Success, "operate success")
	InvalidParamErrno        = New(InvalidParams, "invalid param")
	UnauthorizedErrno        = New(Unauthorized, "unauthorized")
	InternalServerErrorErrno = New(InternalServerError, "internal server failed")
)

// user
var (
	UserRegisterError       = New(UserRegister, "user register failed")
	EncryptPasswordError    = New(EncryptPassword, "encrypt password failed")
	DatabaseCreateUserError = New(DatabaseCreateUser, "create user failed")
	EmailFormatError        = New(EmailFormat, "email format errno")

	UserLoginError               = New(UserLogin, "user login failed")
	GetPasswordFromDatabaseError = New(GetPasswordFromDatabase, "get password digest from db failed")
	CheckPasswordError           = New(CheckPassword, "checkout password failed")
	GenerateTokenError           = New(GenerateToken, "generate token failed")
)
