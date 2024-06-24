package service

import (
	"github.com/stretchr/testify/mock"
	"post-api/internal/model"
)

type MockPostsRepo struct {
	mock.Mock
}

func (m *MockPostsRepo) CreatePost(post model.Post) (uint, error) {
	args := m.Called(post)
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockPostsRepo) GetAllPosts() ([]*model.Post, error) {
	args := m.Called()
	return args.Get(0).([]*model.Post), args.Error(1)
}

func (m *MockPostsRepo) GetPostByID(id uint) (*model.Post, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Post), args.Error(1)
}

func (m *MockPostsRepo) UpdatePost(postID uint, post model.UpdatePostInput) error {
	args := m.Called(post)
	return args.Error(0)
}

type MockCommentsRepo struct {
	mock.Mock
}

func (m *MockCommentsRepo) CreateComment(comment model.Comment) (uint, error) {
	args := m.Called(comment)
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockCommentsRepo) GetCommentsByPostID(postID uint) ([]*model.Comment, error) {
	args := m.Called(postID)
	return args.Get(0).([]*model.Comment), args.Error(1)
}

func (m *MockCommentsRepo) GetCommentsByParentID(parentID uint) ([]*model.Comment, error) {
	args := m.Called(parentID)
	return args.Get(0).([]*model.Comment), args.Error(1)
}
