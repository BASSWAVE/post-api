package model

type Comment struct {
	ID       uint   `json:"id"`
	PostID   uint   `json:"postId"`
	Content  string `json:"content"`
	ParentID *uint  `json:"parentId"`
}
