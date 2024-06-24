package model

type Post struct {
	ID               uint   `json:"id"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	CommentsDisabled bool   `json:"commentsDisabled"`
}
