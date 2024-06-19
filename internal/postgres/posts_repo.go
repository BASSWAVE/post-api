package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"post-api/internal/model"
)

var (
	ErrPostNotFound = errors.New("post not found")
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

func (r *PostsRepository) UpdateCommentsDisabled(postID uint, commentsDisabled bool) (bool, error) {
	_, err := r.pool.Exec(context.Background(),
		`UPDATE posts SET comments_disabled=$1 WHERE id = $2 RETURNING title, content`,
		commentsDisabled, postID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *PostsRepository) ReadAllPosts() ([]*model.Post, error) {
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
			return nil, ErrPostNotFound
		}
		return nil, err
	}
	post.ID = id
	return &post, nil
}

func (r *PostsRepository) GetCommentsStatus(postId uint) (bool, error) {
	var status bool
	err := r.pool.QueryRow(context.Background(),
		`SELECT comments_disabled FROM posts WHERE id=$1`, postId).Scan(&status)
	if err != nil {
		return false, err
	}
	return status, nil
}
