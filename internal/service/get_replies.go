package service

import (
	"post-api/internal/model"
	"strconv"
)

// Извлекаем из репозитория на 1 элемент больше лимита, чтобы проверить наличие элемента на следующей страницу
func (s *Service) GetReplies(commentID uint, limit int, after *string) ([]model.Comment, string, bool, error) {
	if limit < 1 {
		return nil, "", false, nil
	}

	comments, err := s.commentsRepo.GetCommentsByParentID(commentID, limit+1, after)
	if err != nil {
		return nil, "", false, err
	}

	endCursor := ""
	hasNextPage := len(comments) == limit+1
	if hasNextPage {
		endCursor = strconv.Itoa(int(comments[len(comments)-2].ID))
		return comments[:len(comments)-1], endCursor, hasNextPage, nil
	}
	if len(comments) > 0 {
		endCursor = strconv.Itoa(int(comments[len(comments)-1].ID))
	}
	return comments, endCursor, hasNextPage, nil
}
