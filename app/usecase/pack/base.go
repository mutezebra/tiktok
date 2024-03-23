package pack

import (
	"github.com/Mutezebra/tiktok/app/domain/model/errno"
	"github.com/Mutezebra/tiktok/kitex_gen/api/base"
)

var (
	successCode = int32(200)
	successMsg  = "operate success"
	Success     = &base.Base{Code: &successCode, Msg: &successMsg}
)

func NewBase(err errno.Errno) *base.Base {
	msg := err.Error()
	code := int32(err.Code())
	return &base.Base{
		Code: &code,
		Msg:  &msg,
	}
}
