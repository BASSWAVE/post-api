package service

import (
	"post-api/internal/model"
)

type PostsRepo interface {
	CreatePost(post model.PostForCreating) (uint, error)
	GetPosts(limit int, after *string) ([]model.Post, error)
	GetPostByID(id uint) (*model.Post, error)
	UpdatePost(id uint, input model.PostForUpdating) error
}

type CommentsRepo interface {
	CreateComment(comment model.CommentForCreating) (uint, error)
	GetCommentsByPostID(postID uint, limit int, after *string) ([]model.Comment, error)
	GetCommentsByParentID(parentId uint, limit int, after *string) ([]model.Comment, error)
}

type Service struct {
	postsRepo    PostsRepo
	commentsRepo CommentsRepo
}

func NewService(postsRepo PostsRepo, commentsRepo CommentsRepo) *Service {
	return &Service{postsRepo: postsRepo, commentsRepo: commentsRepo}
}
