package service

import "post-api/internal/model"

func (s *Service) GetChildrenComments(parentID uint) ([]model.Comment, error) {
	comments, err := s.commentsRepo.GetCommentsByParentID(parentID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
