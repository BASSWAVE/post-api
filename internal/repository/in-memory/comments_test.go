package in_memory

import (
	"post-api/internal/model"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommentsByParentID(t *testing.T) {
	tests := []struct {
		name           string
		initialStorage map[uint][]model.Comment
		parentID       uint
		expectedResult []model.Comment
	}{
		{
			name: "No comments for parent ID",
			initialStorage: map[uint][]model.Comment{
				1: {{ID: 1, PostID: 1, Content: "First comment"}},
			},
			parentID:       2,
			expectedResult: nil,
		},
		{
			name: "Single comment for parent ID",
			initialStorage: map[uint][]model.Comment{
				1: {{ID: 1, PostID: 1, Content: "First comment"}},
			},
			parentID: 1,
			expectedResult: []model.Comment{
				{ID: 1, PostID: 1, Content: "First comment"},
			},
		},
		{
			name: "Multiple comments for parent ID",
			initialStorage: map[uint][]model.Comment{
				1: {
					{ID: 1, PostID: 1, Content: "First comment"},
					{ID: 2, PostID: 1, Content: "Second comment"},
				},
			},
			parentID: 1,
			expectedResult: []model.Comment{
				{ID: 1, PostID: 1, Content: "First comment"},
				{ID: 2, PostID: 1, Content: "Second comment"},
			},
		},
		{
			name: "Comments for multiple parent IDs",
			initialStorage: map[uint][]model.Comment{
				1: {
					{ID: 1, PostID: 1, Content: "First comment"},
					{ID: 2, PostID: 1, Content: "Second comment"},
				},
				2: {
					{ID: 3, PostID: 1, Content: "Third comment"},
				},
			},
			parentID: 2,
			expectedResult: []model.Comment{
				{ID: 3, PostID: 1, Content: "Third comment"},
			},
		},
		{
			name:           "Empty storage",
			initialStorage: map[uint][]model.Comment{},
			parentID:       1,
			expectedResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &CommentsRepository{
				storageByPostID:   make(map[uint][]model.Comment),
				storageByParentID: tt.initialStorage,
				lastID:            0,
				mx:                &sync.Mutex{},
			}
			result, err := repo.GetCommentsByParentID(tt.parentID)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestGetCommentsByPostID(t *testing.T) {
	tests := []struct {
		name           string
		initialStorage map[uint][]model.Comment
		postID         uint
		expectedResult []model.Comment
	}{
		{
			name: "No comments for post ID",
			initialStorage: map[uint][]model.Comment{
				1: {{ID: 1, PostID: 1, Content: "First comment"}},
			},
			postID:         2,
			expectedResult: nil,
		},
		{
			name: "Single comment for post ID",
			initialStorage: map[uint][]model.Comment{
				1: {{ID: 1, PostID: 1, Content: "First comment"}},
			},
			postID: 1,
			expectedResult: []model.Comment{
				{ID: 1, PostID: 1, Content: "First comment"},
			},
		},
		{
			name: "Multiple comments for post ID",
			initialStorage: map[uint][]model.Comment{
				1: {
					{ID: 1, PostID: 1, Content: "First comment"},
					{ID: 2, PostID: 1, Content: "Second comment"},
				},
			},
			postID: 1,
			expectedResult: []model.Comment{
				{ID: 1, PostID: 1, Content: "First comment"},
				{ID: 2, PostID: 1, Content: "Second comment"},
			},
		},
		{
			name: "Comments for multiple post IDs",
			initialStorage: map[uint][]model.Comment{
				1: {
					{ID: 1, PostID: 1, Content: "First comment"},
					{ID: 2, PostID: 1, Content: "Second comment"},
				},
				2: {
					{ID: 3, PostID: 2, Content: "Third comment"},
				},
			},
			postID: 2,
			expectedResult: []model.Comment{
				{ID: 3, PostID: 2, Content: "Third comment"},
			},
		},
		{
			name:           "Empty storage",
			initialStorage: map[uint][]model.Comment{},
			postID:         1,
			expectedResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &CommentsRepository{
				storageByPostID:   tt.initialStorage,
				storageByParentID: make(map[uint][]model.Comment),
				lastID:            0,
				mx:                &sync.Mutex{},
			}
			result, err := repo.GetCommentsByPostID(tt.postID)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestCreateComment(t *testing.T) {
	tests := []struct {
		name                  string
		initialStoragePost    map[uint][]model.Comment
		initialStorageParent  map[uint][]model.Comment
		commentToCreate       model.CommentForCreating
		lastID                uint
		expectedID            uint
		expectedStoragePost   map[uint][]model.Comment
		expectedStorageParent map[uint][]model.Comment
	}{
		{
			name:                 "Create comment with no parent",
			initialStoragePost:   map[uint][]model.Comment{},
			initialStorageParent: map[uint][]model.Comment{},
			commentToCreate: model.CommentForCreating{
				PostID:  1,
				Content: "First comment",
			},
			lastID:     0,
			expectedID: 1,
			expectedStoragePost: map[uint][]model.Comment{
				1: {{ID: 1, PostID: 1, Content: "First comment"}},
			},
			expectedStorageParent: map[uint][]model.Comment{},
		},
		{
			name:                 "Create comment with parent",
			initialStoragePost:   map[uint][]model.Comment{},
			initialStorageParent: map[uint][]model.Comment{},
			commentToCreate: model.CommentForCreating{
				PostID:    1,
				Content:   "Reply to first comment",
				ParentID:  1,
				HasParent: true,
			},
			expectedID:          1,
			expectedStoragePost: map[uint][]model.Comment{},
			expectedStorageParent: map[uint][]model.Comment{
				1: {{ID: 1, PostID: 1, Content: "Reply to first comment", ParentID: 1, HasParent: true}},
			},
		},
		{
			name:                 "Create multiple comments with no parent",
			initialStoragePost:   map[uint][]model.Comment{},
			initialStorageParent: map[uint][]model.Comment{},
			commentToCreate: model.CommentForCreating{
				PostID:  1,
				Content: "First comment",
			},
			expectedID: 1,
			expectedStoragePost: map[uint][]model.Comment{
				1: {{ID: 1, PostID: 1, Content: "First comment"}},
			},
			expectedStorageParent: map[uint][]model.Comment{},
		},
		{
			name: "Create multiple comments with parent",
			initialStoragePost: map[uint][]model.Comment{
				1: {{ID: 1, PostID: 1, Content: "First comment"}},
			},
			initialStorageParent: map[uint][]model.Comment{},
			lastID:               1,
			commentToCreate: model.CommentForCreating{
				PostID:    1,
				Content:   "Second comment",
				ParentID:  1,
				HasParent: true,
			},
			expectedID: 2,
			expectedStoragePost: map[uint][]model.Comment{
				1: {{ID: 1, PostID: 1, Content: "First comment"}},
			},
			expectedStorageParent: map[uint][]model.Comment{
				1: {
					{ID: 2, PostID: 1, Content: "Second comment", ParentID: 1, HasParent: true},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &CommentsRepository{
				storageByPostID:   tt.initialStoragePost,
				storageByParentID: tt.initialStorageParent,
				lastID:            tt.lastID,
				mx:                &sync.Mutex{},
			}
			id, err := repo.CreateComment(tt.commentToCreate)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedID, id)
			assert.Equal(t, tt.expectedStoragePost, repo.storageByPostID)
			assert.Equal(t, tt.expectedStorageParent, repo.storageByParentID)
		})
	}
}
