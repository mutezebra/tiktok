package errno

var (
	SuccessErrno      = NewErrno(Success, "operate success")
	InvalidParamErrno = NewErrno(InvalidParams, "invalid param")
	UnauthorizedErrno = NewErrno(Unauthorized, "missing authorization")

	InternalServerErrorErrno = NewErrno(InternalServerError, "internal server error")
)
