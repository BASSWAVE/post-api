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

func (r *CommentsRepository) GetCommentsByParentID(parentId uint) ([]model.Comment, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	return r.storageByParentID[parentId], nil
}

func (r *CommentsRepository) GetCommentsByPostID(postID uint) ([]model.Comment, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	return r.storageByPostID[postID], nil
}

func (r *CommentsRepository) CreateComment(comment model.CommentForCreating) (uint, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.lastID++
	commentWithID := model.Comment{
		ID:       r.lastID,
		Content:  comment.Content,
		PostID:   comment.PostID,
		ParentID: comment.ParentID,
	}
	if commentWithID.ParentID != nil {
		r.storageByParentID[*commentWithID.ParentID] = append(r.storageByParentID[*commentWithID.ParentID], commentWithID)
	} else {
		r.storageByPostID[commentWithID.PostID] = append(r.storageByPostID[commentWithID.PostID], commentWithID)
	}
	return commentWithID.ID, nil
}
