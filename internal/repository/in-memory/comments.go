package in_memory

import (
	"post-api/internal/model"
	"sync"
)

type CommentsRepository struct {
	storByPostID   map[uint][]*model.Comment
	storByParentID map[uint][]*model.Comment
	lastID         uint
	mx             sync.Mutex
}

func NewCommentsRepository(storByPostID map[uint][]*model.Comment, storByParentID map[uint][]*model.Comment) *CommentsRepository {
	return &CommentsRepository{storByPostID: storByPostID, storByParentID: storByParentID, lastID: 0}
}

func (r *CommentsRepository) GetCommentsByParentID(parentId uint) ([]*model.Comment, error) {
	r.mx.Lock()
	comment := r.storByParentID[parentId]
	r.mx.Unlock()
	return comment, nil
}

func (r *CommentsRepository) GetCommentsByPostID(postID uint) ([]*model.Comment, error) {
	r.mx.Lock()
	comment := r.storByPostID[postID]
	r.mx.Unlock()
	return comment, nil
}

func (r *CommentsRepository) CreateComment(comment model.Comment) (uint, error) {
	r.lastID++
	comment.ID = r.lastID
	r.mx.Lock()
	r.storByPostID[comment.PostID] = append(r.storByPostID[comment.PostID], &comment)
	r.storByParentID[comment.ParentID] = append(r.storByParentID[comment.ParentID], &comment)
	r.mx.Unlock()
	return r.lastID, nil
}
