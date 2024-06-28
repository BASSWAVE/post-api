package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"post-api/internal/model"
	"post-api/internal/repository"
	"strconv"
	"strings"
)

type PostsRepository struct {
	pool *pgxpool.Pool
}

func NewPostRepository(pool *pgxpool.Pool) *PostsRepository {
	return &PostsRepository{pool: pool}
}

func (r *PostsRepository) CreatePost(post model.PostForCreating) (uint, error) {
	var id uint
	err := r.pool.QueryRow(context.Background(),
		`INSERT INTO posts(title, content, comments_disabled) VALUES ($1, $2, $3) RETURNING id`,
		post.Title, post.Content, post.CommentsDisabled).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PostsRepository) GetPosts(limit int, after *string) ([]model.Post, error) {
	query := `SELECT * FROM posts`
	args := []interface{}{}

	if after != nil {
		afterID, err := strconv.ParseUint(*after, 10, 0)
		if err != nil {
			return nil, err
		}
		query += ` WHERE id > $1`
		args = append(args, afterID)
		query += ` ORDER BY id ASC LIMIT $2`
		args = append(args, limit)
	} else {
		query += ` ORDER BY id ASC LIMIT $1`
		args = append(args, limit)
	}

	rows, err := r.pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Post])
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostsRepository) GetPostByID(id uint) (*model.Post, error) {
	var post model.Post
	err := r.pool.QueryRow(context.Background(),
		`SELECT title, content, comments_disabled FROM posts WHERE id=$1`, id).
		Scan(&post.Title, &post.Content, &post.CommentsDisabled)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrPostNotFound
		}
		return nil, err
	}
	post.ID = id
	return &post, nil
}

func (r *PostsRepository) UpdatePost(postID uint, input model.PostForUpdating) error {
	fields := make([]string, 0)
	args := make([]any, 0)
	argID := 1

	if input.Title != nil {
		fields = append(fields, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}
	if input.Content != nil {
		fields = append(fields, fmt.Sprintf("content=$%d", argID))
		args = append(args, *input.Content)
		argID++
	}
	if input.CommentsDisabled != nil {
		fields = append(fields, fmt.Sprintf("comments_disabled=$%d", argID))
		args = append(args, *input.CommentsDisabled)
		argID++
	}

	if len(fields) == 0 {
		return errors.New("no fields to update")
	}

	query := fmt.Sprintf("UPDATE posts SET %s WHERE id=$%d", strings.Join(fields, ", "), argID)
	args = append(args, postID)

	_, err := r.pool.Exec(context.Background(), query, args...)
	return err
}
