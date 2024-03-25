package oss

import (
	"github.com/Mutezebra/tiktok/config"
	"github.com/qiniu/go-sdk/v7/auth"
)

// conf is the config of oss
var conf config.QiNiu

// InitOSS init the oss config
func InitOSS() {
	conf = *config.Conf.QiNiu
}

// PutRet use to build a message struct that oss upload
// or download needed
type PutRet struct {
	Hash string
	key  string
}

// getOSS return the oss`s max,bucket and domain
func getOSS() (mac *auth.Credentials, bucket string, domain string) {
	acK := conf.AccessKey
	seK := conf.SecretKey
	domain = conf.Domain
	return auth.New(acK, seK), conf.Bucket, domain
}
