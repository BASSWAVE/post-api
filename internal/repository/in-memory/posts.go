package in_memory

import (
	"post-api/internal/model"
	"post-api/internal/repository"
	"sync"
)

type PostsRepository struct {
	storage map[uint]*model.Post
	lastID  uint
	mx      sync.Mutex
}

func NewPostRepository() *PostsRepository {
	return &PostsRepository{storage: make(map[uint]*model.Post)}
}

func (r *PostsRepository) CreatePost(post model.PostForCreating) (uint, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.lastID++
	postWithID := model.Post{
		ID:               r.lastID,
		Title:            post.Title,
		Content:          post.Content,
		CommentsDisabled: post.CommentsDisabled,
	}
	r.storage[postWithID.ID] = &postWithID
	return postWithID.ID, nil
}

func (r *PostsRepository) GetPosts() ([]model.Post, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	posts := make([]model.Post, 0)
	for _, val := range r.storage {
		posts = append(posts, *val)
	}
	return posts, nil
}

func (r *PostsRepository) GetPostByID(id uint) (*model.Post, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	return r.storage[id], nil
}

func (r *PostsRepository) UpdatePost(id uint, input model.PostForUpdating) error {
	r.mx.Lock()
	defer r.mx.Unlock()
	updatedPost := r.storage[id]
	if updatedPost == nil {
		return repository.ErrPostNotFound
	}
	if input.Title != nil {
		updatedPost.Title = *input.Title
	}
	if input.Content != nil {
		updatedPost.Content = *input.Content
	}
	if input.CommentsDisabled != nil {
		updatedPost.CommentsDisabled = *input.CommentsDisabled
	}
	r.storage[id] = updatedPost
	return nil
}
