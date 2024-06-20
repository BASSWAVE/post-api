package service

import "post-api/internal/model"

func (s *Service) GetPost(id uint) (*model.Post, error) {
	post, err := s.postsRepo.GetPostByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}
