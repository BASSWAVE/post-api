package service

import "post-api/internal/model"

func (s *Service) GetPostComments(postID uint) ([]*model.Comment, error) {
	comments, err := s.commentsRepo.GetCommentsByPostID(postID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
