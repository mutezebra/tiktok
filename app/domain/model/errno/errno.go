package errno

type Errno interface {
	Error() string
	Code() int
}

type errno struct {
	code int
	msg  string
}

func (e *errno) Error() string {
	return e.msg
}

func (e *errno) Code() int {
	return e.code
}

func New(code int, msg string) Errno {
	return &errno{
		code: code,
		msg:  msg,
	}
}

type withMessage struct {
	code  int
	msg   string
	cause error
}

func (w *withMessage) Error() string {
	return w.msg + ": " + w.cause.Error()
}

func (w *withMessage) Code() int {
	return w.code
}

func (w *withMessage) Cause() error {
	return w.cause
}

func WithError(errno Errno, err error) Errno {
	if err == nil {
		return errno
	}
	if errno == nil {
		return nil
	}

	return &withMessage{
		code:  errno.Code(),
		msg:   errno.Error(),
		cause: err,
	}

}

func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

func Convert(err error) Errno {
	if err == nil {
		return SuccessErrno
	}
	return &errno{
		code: -1,
		msg:  err.Error(),
	}
}
