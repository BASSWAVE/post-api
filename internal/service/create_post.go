package service

import "post-api/internal/model"

func (s *Service) CreatePost(post model.Post) (uint, error) {
	id, err := s.postsRepo.CreatePost(post)
	if err != nil {
		return 0, err
	}
	return id, nil
}
