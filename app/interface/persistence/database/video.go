package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Mutezebra/tiktok/app/domain/repository"
)

type VideoRepository struct {
	db *sql.DB
}

func NewVideoRepository() *VideoRepository { return &VideoRepository{_db} }

func (repo *VideoRepository) CreateVideo(ctx context.Context, video *repository.Video) (vid int64, err error) {
	res, err := repo.db.ExecContext(ctx, "INSERT INTO video(id,uid, video_url, cover_url, intro, title, stars, favorites, views, create_at, update_at, delete_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)",
		video.ID, video.UID, video.VideoURL, video.CoverURL,
		video.Intro, video.Title, video.Starts, video.Favorites,
		video.Views, video.CreateAt, video.UpdateAt, video.DeleteAt)
	if err != nil {
		return 0, err
	}

	vid, err = res.LastInsertId()
	return vid, err
}
func (repo *VideoRepository) GetVideoInfo(ctx context.Context, vid int64) (*repository.Video, error) {
	var video repository.Video
	err := repo.db.QueryRowContext(ctx, "SELECT * FROM video where id=?", vid).Scan(
		&video.ID, &video.UID, &video.VideoURL, &video.CoverURL,
		&video.Intro, &video.Title, &video.Starts, &video.Favorites,
		&video.Views, &video.CreateAt, &video.UpdateAt, &video.DeleteAt)

	return &video, err
}

func (repo *VideoRepository) GetVideoListByID(ctx context.Context, uid int64, page int, size int) ([]repository.Video, error) {
	offset := (page - 1) * size

	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM video WHERE uid = ? LIMIT ? OFFSET ?", uid, size, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	videos := make([]repository.Video, 0)
	for rows.Next() {
		var video repository.Video
		if err = rows.Scan(&video.ID, &video.UID, &video.VideoURL, &video.CoverURL, &video.Intro, &video.Title, &video.Starts, &video.Favorites, &video.Views, &video.CreateAt, &video.UpdateAt, &video.DeleteAt); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}

func (repo *VideoRepository) GetVideoPopular(ctx context.Context, vids []int64) ([]repository.Video, error) {
	if len(vids) == 0 {
		return nil, nil
	}

	query := "SELECT * FROM video WHERE id IN (?" + strings.Repeat(",?", len(vids)-1) + ")"

	args := make([]interface{}, len(vids))
	for i, v := range vids {
		args[i] = v
	}

	// Execute the query
	rows, err := repo.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	videos := make([]repository.Video, len(vids))
	for rows.Next() {
		var video repository.Video
		if err = rows.Scan(&video.ID, &video.UID, &video.VideoURL, &video.CoverURL, &video.Intro, &video.Title, &video.Starts, &video.Favorites, &video.Views, &video.CreateAt, &video.UpdateAt, &video.DeleteAt); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}
func (repo *VideoRepository) SearchVideo(ctx context.Context, content string, page, size int) ([]repository.Video, error) {
	like := fmt.Sprintf("%%%s%%", content)
	offset := (page - 1) * size

	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM video WHERE intro LIKE ? OR title LIKE ? LIMIT ? offset ?", like, like, size, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	videos := make([]repository.Video, 0)
	for rows.Next() {
		var video repository.Video
		if err = rows.Scan(&video.ID, &video.UID, &video.VideoURL,
			&video.CoverURL, &video.Intro, &video.Title, &video.Starts,
			&video.Favorites, &video.Views, &video.CreateAt, &video.UpdateAt,
			&video.DeleteAt); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}

func (repo *VideoRepository) GetVideoUrl(ctx context.Context, vid int64) (string, error) {
	var url string
	err := repo.db.QueryRowContext(ctx, "SELECT video_url FROM video where id=?", vid).Scan(&url)

	return url, err
}
