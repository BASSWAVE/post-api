package service

import (
	"errors"
	"post-api/internal/model"
)

func (s *Service) CreateComment(comment model.Comment) (uint, error) {
	post, err := s.postsRepo.GetPostByID(comment.PostID)
	if err != nil {
		return 0, err
	}
	if post.CommentsDisabled {
		return 0, errors.New("user has forbidden leaving comments ")
	}

	id, err := s.commentsRepo.CreateComment(comment)
	if err != nil {
		return 0, err
	}

	go func() {
		for _, ch := range model.Subs[post.ID] {
			go func() {
				ch <- &comment
			}()
		}
	}()

	return id, nil
}
