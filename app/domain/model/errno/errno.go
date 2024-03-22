package errno

import "fmt"

type Errno struct {
	code int
	msg  string
}

func (e *Errno) Error() string {
	return fmt.Sprintf("msg:%s", e.msg)
}

func (e *Errno) Code() int {
	return e.code
}

func NewErrno(code int, msg string) *Errno {
	return &Errno{
		code: code,
		msg:  msg,
	}
}

func Convert(err error) *Errno {
	if err == nil {
		return SuccessErrno
	}
	return &Errno{
		code: -1,
		msg:  err.Error(),
	}
}
