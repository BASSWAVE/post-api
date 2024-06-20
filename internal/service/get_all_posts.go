package service

import "post-api/internal/model"

func (s *Service) GetAllPosts() ([]*model.Post, error) {
	posts, err := s.postsRepo.GetAllPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}
