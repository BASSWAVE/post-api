package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"post-api/internal/model"
	"post-api/internal/repository"
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

func (r *PostsRepository) GetAllPosts() ([]*model.Post, error) {
	rows, err := r.pool.Query(context.Background(),
		`SELECT id, title, content, comments_disabled FROM posts`)
	if err != nil {
		return nil, err
	}

	posts, err := pgx.CollectRows(rows, pgx.RowToStructByName[struct {
		ID               uint
		Title            string
		Content          string
		CommentsDisabled bool
	}])
	if err != nil {
		return nil, err
	}
	postsPointers := make([]*model.Post, len(posts))
	for i := range posts {
		postsPointers[i] = &model.Post{
			ID:               posts[i].ID,
			Title:            posts[i].Title,
			Content:          posts[i].Content,
			CommentsDisabled: posts[i].CommentsDisabled,
		}
	}
	return postsPointers, nil
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

func (r *PostsRepository) UpdatePost(post model.Post) error {
	_, err := r.pool.Exec(context.Background(),
		`UPDATE posts SET title=$1, content=$2, comments_disabled=$3 WHERE id=$4`,
		post.Title, post.Content, post.CommentsDisabled, post.ID)
	return err
}
