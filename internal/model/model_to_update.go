package model

type UpdatePostInput struct {
	Title            *string `json:"title"`
	Content          *string `json:"content"`
	CommentsDisabled *bool   `json:"commentsDisabled"`
}
