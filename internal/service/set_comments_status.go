package service

import "post-api/internal/model"

func (s *Service) SetCommentsStatus(postID uint, status bool) error {
	t := new(bool)
	*t = status
	err := s.postsRepo.UpdatePost(postID, model.PostForUpdating{CommentsDisabled: t})
	return err
}
