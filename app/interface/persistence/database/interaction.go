package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Mutezebra/tiktok/app/domain/repository"
)

type InteractionRepository struct {
	db *sql.DB
}

func NewInteractionRepository() *InteractionRepository { return &InteractionRepository{_db} }

// CreateComment create a comment
func (repo *InteractionRepository) CreateComment(ctx context.Context, comment *repository.Comment) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO comment(id,uid, vid,root_id,reply_id,content,likes,create_at,delete_at) VALUES (?,?,?,?,?,?,?,?,?)",
		comment.ID, comment.UID, comment.VID, comment.RootID,
		comment.ReplyID, comment.Content, comment.Likes, comment.CreateAt,
		comment.DeleteAt)
	return err

}

// LikeComment like a comment
func (repo *InteractionRepository) LikeComment(ctx context.Context, uid, cid int64) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE comment SET likes = likes + 1 WHERE id = ?", cid)
	if err != nil {
		return err
	}
	_, err = repo.db.ExecContext(ctx, "INSERT INTO user_comment_likes(user_id,comment_id) VALUES (?,?)", uid, cid)
	return err
}

// DislikeComment dislike a comment
func (repo *InteractionRepository) DislikeComment(ctx context.Context, uid, cid int64) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE comment SET likes = likes - 1 WHERE id = ?", cid)
	if err != nil {
		return err
	}
	_, err = repo.db.ExecContext(ctx, "DELETE FROM user_comment_likes WHERE comment_id = ? AND user_id =? LIMIT 1", cid, uid)
	return err
}

// DeleteComment delete a comment
func (repo *InteractionRepository) DeleteComment(ctx context.Context, uid, cid int64) error {
	var commentUID int64
	err := repo.db.QueryRowContext(ctx, "SELECT uid FROM comment where id=?", cid).Scan(&commentUID)
	if err != nil {
		return err
	}
	if commentUID != uid {
		return errors.New("you are not the owner of this comment")
	}
	_, err = repo.db.ExecContext(ctx, "DELETE FROM comment WHERE id = ?", cid)
	return err
}

// GetCommentRootID get the root comment id of a comment
func (repo *InteractionRepository) GetCommentRootID(ctx context.Context, cid int64) (int64, error) {
	var rootID int64
	err := repo.db.QueryRowContext(ctx, "SELECT root_id FROM comment WHERE id = ?", cid).Scan(&rootID)
	if err != nil {
		return -1, err
	}
	if rootID == 0 {
		rootID = cid
	}
	return rootID, err
}

// GetCommentList get the child comments of a root comment
func (repo *InteractionRepository) GetCommentList(ctx context.Context, cid int64, page int8, size int8) ([]repository.Comment, error) {
	var rootID int64
	err := repo.db.QueryRowContext(ctx, "SELECT root_id FROM comment WHERE id = ? ", cid).Scan(&rootID)
	if err != nil {
		return nil, err
	}
	if rootID != 0 {
		return nil, errors.New("this comment is not a root comment")
	}

	offset := (page - 1) * size
	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM comment WHERE root_id = ? LIMIT ? OFFSET ?", cid, size, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	comments := make([]repository.Comment, 0)
	for rows.Next() {
		var comment repository.Comment
		err = rows.Scan(
			&comment.ID, &comment.UID, &comment.VID, &comment.RootID,
			&comment.ReplyID, &comment.Content, &comment.Likes, &comment.CreateAt,
			&comment.DeleteAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (repo *InteractionRepository) WhetherCommentLikeItemExist(ctx context.Context, uid, cid int64) (bool, error) {
	var exist bool
	err := repo.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM user_comment_likes WHERE user_id = ? AND comment_id = ?)", uid, cid).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (repo *InteractionRepository) WhetherCommentExist(ctx context.Context, cid int64) (bool, error) {
	var exist bool
	err := repo.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM comment WHERE id = ?)", cid).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

// LikeVideo like a video
func (repo *InteractionRepository) LikeVideo(ctx context.Context, uid, vid int64) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE video SET likes = likes + 1 WHERE id = ?", vid)
	if err != nil {
		return err
	}
	_, err = repo.db.ExecContext(ctx, "INSERT INTO user_video_likes(user_id,video_id) VALUES (?,?)", uid, vid)
	return err
}

// DislikeVideo dislike a video
func (repo *InteractionRepository) DislikeVideo(ctx context.Context, uid, vid int64) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE video SET likes = likes - 1 WHERE id = ?", vid)
	if err != nil {
		return err
	}
	_, err = repo.db.ExecContext(ctx, "DELETE FROM user_video_likes WHERE video_id = ? AND user_id =? LIMIT 1", vid, uid)
	return err
}

// LikeList get the user's like list
func (repo *InteractionRepository) LikeList(ctx context.Context, uid int64, page int8, size int8) ([]repository.Video, error) {
	offset := (page - 1) * size
	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM video WHERE id IN (SELECT video_id FROM user_video_likes WHERE user_id = ?) LIMIT ? OFFSET ?", uid, size, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	videos := make([]repository.Video, 0)
	for rows.Next() {
		var video repository.Video
		err = rows.Scan(&video.ID, &video.UID, &video.VideoURL, &video.CoverURL,
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

// GetVideoDirectCommentList get a video's directed comment list
func (repo *InteractionRepository) GetVideoDirectCommentList(ctx context.Context, vid int64, page int8, size int8) ([]repository.Comment, error) {
	offset := (page - 1) * size
	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM comment WHERE vid = ? AND root_id = 0 LIMIT ? OFFSET ?", vid, size, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	comments := make([]repository.Comment, 0)
	for rows.Next() {
		var comment repository.Comment
		err = rows.Scan(&comment.ID, &comment.UID, &comment.VID, &comment.RootID,
			&comment.ReplyID, &comment.Content, &comment.Likes, &comment.CreateAt,
			&comment.DeleteAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (repo *InteractionRepository) WhetherVideoLikeItemExist(ctx context.Context, uid, vid int64) (bool, error) {
	var exist bool
	err := repo.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM user_video_likes WHERE user_id = ? AND video_id = ?)", uid, vid).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (repo *InteractionRepository) WhetherVideoExist(ctx context.Context, vid int64) (bool, error) {
	var exist bool
	err := repo.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM video WHERE id = ?)", vid).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}