package model

type CommentForCreating struct {
	PostID   uint   `db:"post_id"`
	Content  string `db:"content"`
	ParentID *uint  `db:"parent_id"`
}

type PostForCreating struct {
	Title            string `db:"title"`
	Content          string `db:"content"`
	CommentsDisabled bool   `db:"commentsDisabled"`
}

type PostForUpdating struct {
	Title            *string `json:"title"`
	Content          *string `json:"content"`
	CommentsDisabled *bool   `json:"commentsDisabled"`
}
