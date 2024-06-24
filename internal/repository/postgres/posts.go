package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"post-api/internal/model"
	"post-api/internal/repository"
	"strings"
)

type PostsRepository struct {
	pool *pgxpool.Pool
}

func NewPostRepository(pool *pgxpool.Pool) *PostsRepository {
	return &PostsRepository{pool: pool}
}

func (r *PostsRepository) CreatePost(post model.Post) (uint, error) {
	var id uint
	err := r.pool.QueryRow(context.Background(),
		`INSERT INTO posts(title, content, comments_disabled) VALUES ($1, $2, $3) RETURNING id`,
		post.Title, post.Content, post.CommentsDisabled).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PostsRepository) GetAllPosts() ([]model.Post, error) {
	rows, err := r.pool.Query(context.Background(), `SELECT * FROM posts`)
	if err != nil {
		return nil, err
	}
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

func (r *PostsRepository) UpdatePost(postID uint, post model.UpdatePostInput) error {
	fields := make([]string, 0)
	args := make([]any, 0)
	argID := 1

	if post.Title != nil {
		fields = append(fields, fmt.Sprintf("title=$%d", argID))
		args = append(args, *post.Title)
		argID++
	}
	if post.Content != nil {
		fields = append(fields, fmt.Sprintf("content=$%d", argID))
		args = append(args, *post.Content)
		argID++
	}
	if post.CommentsDisabled != nil {
		fields = append(fields, fmt.Sprintf("comments_disabled=$%d", argID))
		args = append(args, *post.CommentsDisabled)
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
