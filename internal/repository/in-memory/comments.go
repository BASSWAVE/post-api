package in_memory

import (
	"post-api/internal/model"
	"sync"
)

type CommentsRepository struct {
	storageByID       map[uint]*model.Comment
	storageByPostID   map[uint][]*model.Comment
	storageByParentID map[uint][]*model.Comment
	lastID            uint
	mx                *sync.Mutex
}

func NewCommentsRepository() *CommentsRepository {
	return &CommentsRepository{
		storageByID:       make(map[uint]*model.Comment),
		storageByPostID:   make(map[uint][]*model.Comment),
		storageByParentID: make(map[uint][]*model.Comment),
		lastID:            0,
		mx:                &sync.Mutex{},
	}
}

func (r *CommentsRepository) GetCommentByID(id uint) (model.Comment, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	comment := r.storageByID[id]
	return *comment, nil
}

func (r *CommentsRepository) GetCommentsByParentID(parentId uint) ([]model.Comment, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	commentsPointers := r.storageByParentID[parentId]
	comments := make([]model.Comment, len(commentsPointers))
	for i := range commentsPointers {
		comments[i] = *commentsPointers[i]
	}
	return comments, nil
}

func (r *CommentsRepository) GetCommentsByPostID(postID uint) ([]model.Comment, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	commentsPointers := r.storageByPostID[postID]
	comments := make([]model.Comment, len(commentsPointers))
	for i := range commentsPointers {
		comments[i] = *commentsPointers[i]
	}
	return comments, nil
}

func (r *CommentsRepository) CreateComment(comment model.CommentFromServiceToRepo) (uint, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.lastID++
	commentWithID := model.Comment{
		ID:       r.lastID,
		Content:  comment.Content,
		PostID:   comment.PostID,
		ParentID: comment.ParentID,
	}
	r.storageByID[commentWithID.ID] = &commentWithID
	if commentWithID.ParentID != nil {
		r.storageByParentID[*commentWithID.ParentID] = append(r.storageByParentID[*commentWithID.ParentID], &commentWithID)
	} else {
		r.storageByPostID[commentWithID.PostID] = append(r.storageByPostID[commentWithID.PostID], &commentWithID)
	}
	return commentWithID.ID, nil
}
