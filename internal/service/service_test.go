package service_test

import (
	"github.com/stretchr/testify/assert"
	"post-api/internal/model"
	"post-api/internal/service"
	"testing"
)

func TestSetCommentsStatus(t *testing.T) {
	t.Run("set status to an existing post", func(t *testing.T) {
		//arrange
		mockPostsRepo := new(service.MockPostsRepo)
		mockCommentsRepo := new(service.MockCommentsRepo)
		svc := service.NewService(mockPostsRepo, mockCommentsRepo)

		post := model.Post{ID: 1, CommentsDisabled: false}
		updatedPost := model.Post{ID: 1, CommentsDisabled: true}

		mockPostsRepo.On("GetPostByID", uint(1)).Return(&post, nil)
		mockPostsRepo.On("UpdatePost", updatedPost).Return(nil)

		//act
		err := svc.SetCommentsStatus(1, true)

		//assert
		assert.NoError(t, err)

		mockPostsRepo.AssertExpectations(t)
	})
	t.Run("set status to not existing post", func(t *testing.T) {
		//arrange
		mockPostsRepo := new(service.MockPostsRepo)

		post := &model.Post{ID: 1, CommentsDisabled: false}
		newPost := &model.Post{ID: 1, CommentsDisabled: true}

		mockPostsRepo.On("GetPostByID", uint(1)).Return(post, nil)
		mockPostsRepo.On("UpdatePost", *newPost).Return(nil)

		//act
		//assert

	})
}

func TestGetPostComments(t *testing.T) {
	mockPostsRepo := new(service.MockPostsRepo)
	mockCommentsRepo := new(service.MockCommentsRepo)
	svc := service.NewService(mockPostsRepo, mockCommentsRepo)

	comments := []*model.Comment{{ID: 1, PostID: 1}, {ID: 2, PostID: 1}}

	mockCommentsRepo.On("GetCommentsByPostID", uint(1)).Return(comments, nil)

	result, err := svc.GetPostComments(1)
	assert.NoError(t, err)
	assert.Equal(t, comments, result)

	mockCommentsRepo.AssertExpectations(t)
}

func TestCreatePost(t *testing.T) {
	mockPostsRepo := new(service.MockPostsRepo)
	mockCommentsRepo := new(service.MockCommentsRepo)
	svc := service.NewService(mockPostsRepo, mockCommentsRepo)

	post := model.Post{Title: "Test Post"}
	mockPostsRepo.On("CreatePost", post).Return(uint(1), nil)

	id, err := svc.CreatePost(post)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), id)

	mockPostsRepo.AssertExpectations(t)
}

func TestCreateComment(t *testing.T) {
	mockPostsRepo := new(service.MockPostsRepo)
	mockCommentsRepo := new(service.MockCommentsRepo)
	svc := service.NewService(mockPostsRepo, mockCommentsRepo)

	post := &model.Post{ID: 1, CommentsDisabled: false}
	comment := model.Comment{PostID: 1, Content: "Test Comment"}

	mockPostsRepo.On("GetPostByID", uint(1)).Return(post, nil)
	mockCommentsRepo.On("CreateComment", comment).Return(uint(1), nil)

	id, err := svc.CreateComment(comment)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), id)

	mockPostsRepo.AssertExpectations(t)
	mockCommentsRepo.AssertExpectations(t)
}

func TestCreateComment_CommentsDisabled(t *testing.T) {
	mockPostsRepo := new(service.MockPostsRepo)
	mockCommentsRepo := new(service.MockCommentsRepo)
	svc := service.NewService(mockPostsRepo, mockCommentsRepo)

	post := &model.Post{ID: 1, CommentsDisabled: true}
	comment := model.Comment{PostID: 1, Content: "Test Comment"}

	mockPostsRepo.On("GetPostByID", uint(1)).Return(post, nil)

	id, err := svc.CreateComment(comment)
	assert.Error(t, err)
	assert.Equal(t, uint(0), id)
	assert.Equal(t, "user has forbidden leaving comments ", err.Error())

	mockPostsRepo.AssertExpectations(t)
}
