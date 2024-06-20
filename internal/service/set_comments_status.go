package service

func (s *Service) SetCommentsStatus(postID uint, status bool) error {
	post, err := s.postsRepo.GetPostByID(postID)
	if err != nil {
		return err
	}
	if post.CommentsDisabled == status {
		return nil
	}
	post.CommentsDisabled = status
	err = s.postsRepo.UpdatePost(*post)
	return err
}
