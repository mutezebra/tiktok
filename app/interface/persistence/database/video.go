package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Mutezebra/tiktok/pkg/log"

	"github.com/Mutezebra/tiktok/app/domain/repository"
)

type VideoRepository struct {
	db *sql.DB
}

func NewVideoRepository() *VideoRepository { return &VideoRepository{_db} }

func (repo *VideoRepository) CreateVideo(ctx context.Context, video *repository.Video) (vid int64, err error) {
	res, err := repo.db.ExecContext(ctx, "INSERT INTO video(id,uid, video_url, cover_url, intro, title, video_ext,cover_ext,stars, likes, views, create_at, update_at, delete_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		video.ID, video.UID, video.VideoURL, video.CoverURL,
		video.Intro, video.Title, video.VideoExt, video.CoverExt, video.Starts,
		video.Likes, video.Views, video.CreateAt, video.UpdateAt,
		video.DeleteAt)
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
		&video.Intro, &video.Title, &video.VideoExt, &video.CoverExt, &video.Starts,
		&video.Likes, &video.Views, &video.CreateAt, &video.UpdateAt,
		&video.DeleteAt)

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
		err = rows.Scan(
			&video.ID, &video.UID, &video.VideoURL, &video.CoverURL,
			&video.Intro, &video.Title, &video.VideoExt, &video.CoverExt, &video.Starts,
			&video.Likes, &video.Views, &video.CreateAt, &video.UpdateAt,
			&video.DeleteAt)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}

func (repo *VideoRepository) GetVideosInfo(ctx context.Context, vids []int64) ([]repository.Video, error) {
	stmt, err := repo.db.Prepare("SELECT * FROM video WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer func() { _ = stmt.Close() }()

	videos := make([]repository.Video, 0, len(vids))

	for _, vid := range vids {
		row := stmt.QueryRowContext(ctx, vid)
		if row.Err() != nil {
			return nil, row.Err()
		}
		video := repository.Video{}
		if err = row.Scan(
			&video.ID, &video.UID, &video.VideoURL, &video.CoverURL,
			&video.Intro, &video.Title, &video.VideoExt, &video.CoverExt, &video.Starts,
			&video.Likes, &video.Views, &video.CreateAt, &video.UpdateAt,
			&video.DeleteAt); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func (repo *VideoRepository) SearchVideo(ctx context.Context, content string, page int, size int) ([]repository.Video, error) {
	like := fmt.Sprintf("%%%s%%", content)
	offset := (page - 1) * size

	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM video WHERE intro LIKE ? OR title LIKE ? LIMIT ? OFFSET ?", like, like, size, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	videos := make([]repository.Video, 0)
	for rows.Next() {
		var video repository.Video
		if err = rows.Scan(
			&video.ID, &video.UID, &video.VideoURL, &video.CoverURL,
			&video.Intro, &video.Title, &video.VideoExt, &video.CoverExt, &video.Starts,
			&video.Likes, &video.Views, &video.CreateAt, &video.UpdateAt,
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
	err := repo.db.QueryRowContext(ctx, "SELECT video_url FROM video where id=? LIMIT 1", vid).Scan(&url)

	return url, err
}

func (repo *VideoRepository) GetValByColumn(ctx context.Context, vid int64, column string) (string, error) {
	var val string
	query := fmt.Sprintf("SELECT %s FROM video WHERE id=%d LIMIT 1", column, vid)
	err := repo.db.QueryRowContext(ctx, query).Scan(&val)
	return val, err
}

func (repo *VideoRepository) GetVideoViews(ctx context.Context, vid int64) (int32, error) {
	var views int32
	err := repo.db.QueryRowContext(ctx, "SELECT views FROM video where id=? LIMIT 1", vid).Scan(&views)

	return views, err
}

func (repo *VideoRepository) UpdateViews(kvs map[int64]int32) {
	stmt, err := repo.db.Prepare("UPDATE video SET views=? WHERE id=?")
	if err != nil {
		log.LogrusObj.Panic(fmt.Sprintf("create prepare sql failed,cause:%s", err.Error()))
	}
	defer func() { _ = stmt.Close() }()

	for id, views := range kvs {
		_, err = stmt.Exec(views, id)
		if err != nil {
			log.LogrusObj.Warning(fmt.Sprintf("update views to database failed,cause:%s", err.Error()))
			return
		}
	}
}
