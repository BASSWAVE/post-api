package service

import (
	"post-api/internal/model"
	"strconv"
)

// Извлекаем из репозитория на 1 элемент больше лимита, чтобы проверить наличие элемента на следующей страницу
func (s *Service) GetPosts(limit int, after *string) ([]model.Post, string, bool, error) {
	if limit < 1 {
		return nil, "", false, nil
	}

	posts, err := s.postsRepo.GetPosts(limit+1, after)
	if err != nil {
		return nil, "", false, err
	}

	endCursor := ""
	hasNextPage := len(posts) == limit+1
	if hasNextPage {
		endCursor = strconv.Itoa(int(posts[len(posts)-2].ID))
		return posts[:len(posts)-1], endCursor, hasNextPage, nil
	}
	if len(posts) > 0 {
		endCursor = strconv.Itoa(int(posts[len(posts)-1].ID))
	}
	return posts, endCursor, hasNextPage, nil
}
