package model

type CommentFromServiceToRepo struct {
	PostID   uint   `db:"post_id"`
	Content  string `db:"content"`
	ParentID *uint  `db:"parent_id"`
}

type PostFromServiceToRepo struct {
	Title            string `db:"title"`
	Content          string `db:"content"`
	CommentsDisabled bool   `db:"commentsDisabled"`
}
