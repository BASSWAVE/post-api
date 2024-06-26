package in_memory

import (
	"github.com/stretchr/testify/assert"
	"post-api/internal/model"
	"post-api/internal/repository"
	"sync"
	"testing"
)

func TestCreatePost(t *testing.T) {
	tests := []struct {
		name            string
		initialStorage  map[uint]*model.Post
		postToCreate    model.PostForCreating
		expectedID      uint
		expectedStorage map[uint]*model.Post
	}{
		{
			name:           "Create first post",
			initialStorage: map[uint]*model.Post{},
			postToCreate: model.PostForCreating{
				Title:            "First post",
				Content:          "This is the first post.",
				CommentsDisabled: false,
			},
			expectedID: 1,
			expectedStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "First post",
					Content:          "This is the first post.",
					CommentsDisabled: false,
				},
			},
		},
		{
			name:           "Create post with comments disabled",
			initialStorage: map[uint]*model.Post{},
			postToCreate: model.PostForCreating{
				Title:            "Second post",
				Content:          "This is the second post.",
				CommentsDisabled: true,
			},
			expectedID: 1,
			expectedStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "Second post",
					Content:          "This is the second post.",
					CommentsDisabled: true,
				},
			},
		},
		{
			name: "Create multiple posts",
			initialStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "First post",
					Content:          "This is the first post.",
					CommentsDisabled: false,
				},
			},
			postToCreate: model.PostForCreating{
				Title:            "Second post",
				Content:          "This is the second post.",
				CommentsDisabled: false,
			},
			expectedID: 2,
			expectedStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "First post",
					Content:          "This is the first post.",
					CommentsDisabled: false,
				},
				2: {
					ID:               2,
					Title:            "Second post",
					Content:          "This is the second post.",
					CommentsDisabled: false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &PostsRepository{
				storage: tt.initialStorage,
				lastID:  uint(len(tt.initialStorage)),
				mx:      sync.Mutex{},
			}
			id, err := repo.CreatePost(tt.postToCreate)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedID, id)
			assert.Equal(t, tt.expectedStorage, repo.storage)
		})
	}
}

func TestGetAllPosts(t *testing.T) {
	tests := []struct {
		name           string
		initialStorage map[uint]*model.Post
		expectedResult []model.Post
	}{
		{
			name:           "No posts in storage",
			initialStorage: map[uint]*model.Post{},
			expectedResult: []model.Post{},
		},
		{
			name: "Single post in storage",
			initialStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "First post",
					Content:          "This is the first post.",
					CommentsDisabled: false,
				},
			},
			expectedResult: []model.Post{
				{
					ID:               1,
					Title:            "First post",
					Content:          "This is the first post.",
					CommentsDisabled: false,
				},
			},
		},
		{
			name: "Multiple posts in storage",
			initialStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "First post",
					Content:          "This is the first post.",
					CommentsDisabled: false,
				},
				2: {
					ID:               2,
					Title:            "Second post",
					Content:          "This is the second post.",
					CommentsDisabled: true,
				},
			},
			expectedResult: []model.Post{
				{
					ID:               1,
					Title:            "First post",
					Content:          "This is the first post.",
					CommentsDisabled: false,
				},
				{
					ID:               2,
					Title:            "Second post",
					Content:          "This is the second post.",
					CommentsDisabled: true,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &PostsRepository{
				storage: tt.initialStorage,
				lastID:  uint(len(tt.initialStorage)),
				mx:      sync.Mutex{},
			}
			result, err := repo.GetAllPosts()
			assert.NoError(t, err)
			assert.ElementsMatch(t, tt.expectedResult, result)
		})
	}
}

func TestGetPostByID(t *testing.T) {
	tests := []struct {
		name           string
		initialStorage map[uint]*model.Post
		postID         uint
		expectedResult *model.Post
	}{
		{
			name:           "Post not found",
			initialStorage: map[uint]*model.Post{},
			postID:         1,
			expectedResult: nil,
		},
		{
			name: "Single post found",
			initialStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "First post",
					Content:          "This is the first post.",
					CommentsDisabled: false,
				},
			},
			postID: 1,
			expectedResult: &model.Post{
				ID:               1,
				Title:            "First post",
				Content:          "This is the first post.",
				CommentsDisabled: false,
			},
		},
		{
			name: "Multiple posts, get second",
			initialStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "First post",
					Content:          "This is the first post.",
					CommentsDisabled: false,
				},
				2: {
					ID:               2,
					Title:            "Second post",
					Content:          "This is the second post.",
					CommentsDisabled: true,
				},
			},
			postID: 2,
			expectedResult: &model.Post{
				ID:               2,
				Title:            "Second post",
				Content:          "This is the second post.",
				CommentsDisabled: true,
			},
		},
		{
			name: "Post not found in non-empty storage",
			initialStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "First post",
					Content:          "This is the first post.",
					CommentsDisabled: false,
				},
			},
			postID:         2,
			expectedResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &PostsRepository{
				storage: tt.initialStorage,
				lastID:  uint(len(tt.initialStorage)),
				mx:      sync.Mutex{},
			}
			result, err := repo.GetPostByID(tt.postID)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestUpdatePost(t *testing.T) {
	title1 := "Updated Title"
	content1 := "Updated Content"
	commentsDisabled1 := true

	tests := []struct {
		name            string
		initialStorage  map[uint]*model.Post
		postID          uint
		input           model.PostForUpdating
		expectedError   error
		expectedStorage map[uint]*model.Post
	}{
		{
			name: "Update title and content",
			initialStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "Original Title",
					Content:          "Original Content",
					CommentsDisabled: false,
				},
			},
			postID: 1,
			input: model.PostForUpdating{
				Title:   &title1,
				Content: &content1,
			},
			expectedError: nil,
			expectedStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "Updated Title",
					Content:          "Updated Content",
					CommentsDisabled: false,
				},
			},
		},
		{
			name: "Update comments disabled",
			initialStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "Original Title",
					Content:          "Original Content",
					CommentsDisabled: false,
				},
			},
			postID: 1,
			input: model.PostForUpdating{
				CommentsDisabled: &commentsDisabled1,
			},
			expectedError: nil,
			expectedStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "Original Title",
					Content:          "Original Content",
					CommentsDisabled: true,
				},
			},
		},
		{
			name:           "Post not found",
			initialStorage: map[uint]*model.Post{},
			postID:         1,
			input: model.PostForUpdating{
				Title: &title1,
			},
			expectedError:   repository.ErrPostNotFound,
			expectedStorage: map[uint]*model.Post{},
		},
		{
			name: "No updates",
			initialStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "Original Title",
					Content:          "Original Content",
					CommentsDisabled: false,
				},
			},
			postID:        1,
			input:         model.PostForUpdating{},
			expectedError: nil,
			expectedStorage: map[uint]*model.Post{
				1: {
					ID:               1,
					Title:            "Original Title",
					Content:          "Original Content",
					CommentsDisabled: false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &PostsRepository{
				storage: tt.initialStorage,
				lastID:  uint(len(tt.initialStorage)),
				mx:      sync.Mutex{},
			}
			err := repo.UpdatePost(tt.postID, tt.input)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedStorage, repo.storage)
		})
	}
}
