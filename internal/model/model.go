package model

type Comment struct {
	ID       uint   `json:"id"`
	PostID   uint   `json:"postId"`
	Content  string `json:"content"`
	ParentID *uint  `json:"parentId"`
}

type Post struct {
	ID               uint   `json:"id"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	CommentsDisabled bool   `json:"commentsDisabled"`
}
