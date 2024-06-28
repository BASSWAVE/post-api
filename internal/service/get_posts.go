package service

import (
	"post-api/internal/model"
	"strconv"
)

func (s *Service) GetPosts(limit int, after *string) ([]model.Post, string, bool, error) {
	posts, err := s.postsRepo.GetPosts(limit, after)
	if err != nil {
		return nil, "", false, err
	}

	// Determine if there's a next page
	hasNextPage := len(posts) == limit

	// Set end cursor
	endCursor := ""
	if len(posts) > 0 {
		endCursor = strconv.Itoa(int(posts[len(posts)-1].ID))
	}

	return posts, endCursor, hasNextPage, nil
}
