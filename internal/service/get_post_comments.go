package service

import (
	"post-api/internal/model"
	"strconv"
)

func (s *Service) GetPostComments(postID uint, limit int, after *string) ([]model.Comment, string, bool, error) {
	comments, err := s.commentsRepo.GetCommentsByPostID(postID, limit, after)
	if err != nil {
		return nil, "", false, err
	}

	hasNextPage := len(comments) == limit

	endCursor := ""
	if len(comments) > 0 {
		endCursor = strconv.Itoa(int(comments[len(comments)-1].ID))
	}

	return comments, endCursor, hasNextPage, nil
}
