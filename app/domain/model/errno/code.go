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
)
