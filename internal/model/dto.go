package model

type CommentForCreating struct {
	PostID    uint
	Content   string
	ParentID  uint
	HasParent bool
}

type PostForCreating struct {
	Title            string
	Content          string
	CommentsDisabled bool
}

type PostForUpdating struct {
	Title            *string
	Content          *string
	CommentsDisabled *bool
}

func BuildPost(id uint, postWithoutID PostForCreating) Post {
	return Post{
		ID:               id,
		Title:            postWithoutID.Title,
		Content:          postWithoutID.Content,
		CommentsDisabled: postWithoutID.CommentsDisabled,
	}
}

func BuildComment(id uint, commentWithoutID CommentForCreating) Comment {
	return Comment{
		ID:        id,
		PostID:    commentWithoutID.PostID,
		Content:   commentWithoutID.Content,
		ParentID:  commentWithoutID.ParentID,
		HasParent: commentWithoutID.HasParent,
	}
}
