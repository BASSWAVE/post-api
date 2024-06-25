package service

import "post-api/internal/model"

func (s *Service) GetReplies(commentID uint) ([]model.Comment, error) {
	comments, err := s.commentsRepo.GetCommentsByParentID(commentID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
