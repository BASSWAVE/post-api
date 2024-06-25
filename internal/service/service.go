package service

import (
	"post-api/internal/model"
)

type PostsRepo interface {
	CreatePost(post model.PostForCreating) (uint, error)
	GetAllPosts() ([]model.Post, error)
	GetPostByID(id uint) (*model.Post, error)
	UpdatePost(id uint, input model.PostForUpdating) error
}

type CommentsRepo interface {
	CreateComment(comment model.CommentForCreating) (uint, error)
	GetCommentsByPostID(postID uint) ([]model.Comment, error)
	GetCommentsByParentID(parentId uint) ([]model.Comment, error)
}

type Service struct {
	postsRepo    PostsRepo
	commentsRepo CommentsRepo
}

func NewService(postsRepo PostsRepo, commentsRepo CommentsRepo) *Service {
	return &Service{postsRepo: postsRepo, commentsRepo: commentsRepo}
}
