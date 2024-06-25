package model

type Comment struct {
	ID        uint
	PostID    uint
	ParentID  uint
	Content   string
	HasParent bool
}

type Post struct {
	ID               uint
	Title            string
	Content          string
	CommentsDisabled bool
}
