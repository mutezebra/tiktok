package consts

import "time"

const (
	CacheVideoViewKeyPrefix          = "video:view:"
	CacheVideoFavoriteKeyPrefix      = "video:favorites:"
	CacheVideoPopularKey             = "video:popular:primary"
	CacheVideoPopularVideosSize      = 50
	CacheVideoPopularRefreshInterval = 5 * time.Second
	CacheVideoViewExpireTime         = 10 * time.Second
	CacheVideoFavoriteExpireTime     = 60 * time.Second

	DatabaseDefaultUpdateViewInterval = 5 * time.Second
)
