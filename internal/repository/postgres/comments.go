package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"post-api/internal/model"
	"strconv"
)

type CommentsRepository struct {
	pool *pgxpool.Pool
}

func NewCommentsRepository(pool *pgxpool.Pool) *CommentsRepository {
	return &CommentsRepository{pool: pool}
}

func (r *CommentsRepository) GetCommentsByParentID(parentId uint, limit int, after *string) ([]model.Comment, error) {
	query := `SELECT * FROM comments WHERE parent_id = $1`
	args := []interface{}{parentId}

	if after != nil {
		afterID, err := strconv.ParseUint(*after, 10, 0)
		if err != nil {
			return nil, err
		}
		query += ` AND id > $2`
		args = append(args, afterID)
		query += ` ORDER BY id ASC LIMIT $3`
		args = append(args, limit)
	} else {
		query += ` ORDER BY id ASC LIMIT $2`
		args = append(args, limit)
	}

	rows, err := r.pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Comment])
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentsRepository) GetCommentsByPostID(postID uint, limit int, after *string) ([]model.Comment, error) {
	query := `SELECT * FROM comments WHERE post_id = $1 AND has_parent IS false`
	args := []interface{}{postID}

	if after != nil {
		afterID, err := strconv.ParseUint(*after, 10, 0)
		if err != nil {
			return nil, err
		}
		query += ` AND id > $2`
		args = append(args, afterID)
		query += ` ORDER BY id ASC LIMIT $3`
		args = append(args, limit)
	} else {
		query += ` ORDER BY id ASC LIMIT $2`
		args = append(args, limit)
	}

	rows, err := r.pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
