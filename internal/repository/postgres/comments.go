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
		`SELECT * FROM comments WHERE post_id = $1 AND parent_id IS NULL`, postID)
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
		`INSERT INTO comments(post_id, content, parent_id) VALUES ($1, $2, $3) RETURNING id`,
		comment.PostID, comment.Content, comment.ParentID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
