package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"post-api/internal/model"
)

type CommentsRepository struct {
	pool *pgxpool.Pool
}

func NewCommentsRepository(pool *pgxpool.Pool) *CommentsRepository {
	return &CommentsRepository{pool: pool}
}

func (r *CommentsRepository) GetCommentsByParentID(parentId uint) ([]model.Comment, error) {
	rows, err := r.pool.Query(context.Background(), `SELECT * FROM comments WHERE parent_id = $1`, parentId)
	if err != nil {
		return nil, err
	}
	comments, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Comment])
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentsRepository) GetCommentsByPostID(postID uint) ([]model.Comment, error) {
	rows, err := r.pool.Query(context.Background(),
		`SELECT * FROM comments WHERE post_id = $1 AND has_parent IS false`, postID)
	if err != nil {
		return nil, err
	}
	comments, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Comment])
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentsRepository) CreateComment(comment model.CommentForCreating) (uint, error) {
	var id uint
	err := r.pool.QueryRow(context.Background(),
		`INSERT INTO comments(post_id, content, parent_id, has_parent) VALUES ($1, $2, $3, $4) RETURNING id`,
		comment.PostID, comment.Content, comment.ParentID, comment.HasParent).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
