package service

import (
	"post-api/internal/model"
	"strconv"
)

func (s *Service) GetReplies(commentID uint, limit int, after *string) ([]model.Comment, string, bool, error) {
	comments, err := s.commentsRepo.GetCommentsByParentID(commentID, limit, after)
	if err != nil {
		return nil, "", false, err
	}

	// Determine if there's a next page
	hasNextPage := len(comments) == limit

	// Set end cursor
	endCursor := ""
	if len(comments) > 0 {
		endCursor = strconv.Itoa(int(comments[len(comments)-1].ID))
	}

	return comments, endCursor, hasNextPage, nil
}
