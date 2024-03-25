package pack

import (
	errno2 "github.com/Mutezebra/tiktok/app/domain/model/errno"
	"github.com/Mutezebra/tiktok/pkg/log"
)

func ReturnError(errno errno2.Errno, err error) error {
	errno = errno2.WithError(errno, err)
	log.LogrusObj.Error(errno)
	return errno
}
