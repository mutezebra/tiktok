package pack

import (
	"errors"
	"net/http"

	"github.com/Mutezebra/tiktok/pkg/log"

	"github.com/Mutezebra/tiktok/app/domain/model/errno"

	"github.com/cloudwego/hertz/pkg/app"
)

type Response struct {
	Base Base `json:"base"`
	Data any  `json:"data,omitempty"`
}

type Base struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func SendResponse(c *app.RequestContext, data any) {
	resp := &Response{
		Base: Base{
			200,
			"operate success",
		},
	}
	if data != nil {
		resp.Data = data
	}

	c.JSON(http.StatusOK, resp)
}

func SendFailedResponse(c *app.RequestContext, err error) {
	log.LogrusObj.Error(err)
	var e errno.Errno
	ok := errors.As(err, &e)
	if !ok {
		e = errno.Convert(err)
	}

	resp := &Response{
		Base: Base{
			e.BizStatusCode(),
			e.BizMessage(),
		},
	}
	c.JSON(http.StatusOK, resp)
}
