package resolver

import "post-api/internal/postgres"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type PostsRepository interface {
}

type Resolver struct {
	postsRepo    *postgres.PostsRepository
	commentsRepo *postgres.CommentsRepository
}

func NewResolver(postsRepo *postgres.PostsRepository, commentsRepo *postgres.CommentsRepository) *Resolver {
	return &Resolver{
		postsRepo:    postsRepo,
		commentsRepo: commentsRepo,
	}
}
