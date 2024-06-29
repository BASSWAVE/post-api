package in_memory

import (
	"post-api/internal/model"
	"sync"
)

type CommentsRepository struct {
	storageByPostID   map[uint][]model.Comment
	storageByParentID map[uint][]model.Comment
	lastID            uint
	mx                *sync.Mutex
}

func NewCommentsRepository() *CommentsRepository {
	return &CommentsRepository{
		storageByPostID:   make(map[uint][]model.Comment),
		storageByParentID: make(map[uint][]model.Comment),
		lastID:            0,
		mx:                &sync.Mutex{},
	}
}

func (r *CommentsRepository) GetCommentsByParentID(parentId uint, limit int, after *string) ([]model.Comment, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	return r.storageByParentID[parentId], nil
}

func (r *CommentsRepository) GetCommentsByPostID(postID uint, limit int, after *string) ([]model.Comment, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	return r.storageByPostID[postID], nil
}

func (r *CommentsRepository) CreateComment(comment model.CommentForCreating) (uint, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.lastID++
	commentToAdd := model.BuildComment(r.lastID, comment)
	if commentToAdd.HasParent {
		r.storageByParentID[commentToAdd.ParentID] = append(r.storageByParentID[commentToAdd.ParentID], commentToAdd)
	} else {
		r.storageByPostID[commentToAdd.PostID] = append(r.storageByPostID[commentToAdd.PostID], commentToAdd)
	}
	return commentToAdd.ID, nil
}
