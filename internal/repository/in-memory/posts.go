package in_memory

import (
	"post-api/internal/model"
	"sync"
)

type PostsRepository struct {
	storage map[uint]model.PostFromServiceToRepo
	lastID  uint
	mx      sync.Mutex
}

func NewPostRepository(storage map[uint]*model.Post) *PostsRepository {
	return &PostsRepository{storage: storage}
}

func (r *PostsRepository) CreatePost(post model.Post) (uint, error) {
	r.lastID++
	post.ID = r.lastID
	r.mx.Lock()
	r.storage[post.ID] = &post
	r.mx.Unlock()
	return post.ID, nil
}

func (r *PostsRepository) GetAllPosts() ([]*model.Post, error) {
	posts := make([]*model.Post, 0)
	r.mx.Lock()
	for _, val := range r.storage {
		posts = append(posts, val)
	}
	r.mx.Unlock()
	return posts, nil
}

func (r *PostsRepository) GetPostByID(id uint) (*model.Post, error) {
	r.mx.Lock()
	post := r.storage[id]
	r.mx.Unlock()
	return post, nil
}

func (r *PostsRepository) UpdatePost(post model.Post) error {
	r.mx.Lock()
	r.storage[post.ID] = &post
	r.mx.Unlock()
	return nil
}
