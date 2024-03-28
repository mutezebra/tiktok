package repository

type VideoCacheRepository interface {
	GetPopularVids() ([]int64, error)
	StoreVideoViews(vid int64)
	IncrVideoViews(vit int64)
	StoreVideoFavorites(vid int64)
	IncrVideoFavorites(vid int64)
}
