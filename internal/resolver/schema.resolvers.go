package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"log"
	"post-api/internal/graph"
	"post-api/internal/model"
	"strconv"
)

const defaultLimit = 10

// Replies is the resolver for the replies field.
func (r *commentResolver) Replies(ctx context.Context, obj *model.Comment, first *int, after *string) (*model.CommentConnection, error) {
	log.Println("comment resolver: replies")

	limit := defaultLimit
	if first != nil {
		limit = *first
	}

	replies, endCursor, hasNextPage, err := r.serv.GetReplies(obj.ID, limit, after)
	if err != nil {
		return nil, err
	}

	var edges []model.CommentEdge
	for _, reply := range replies {
		edges = append(edges, model.CommentEdge{
			Cursor: strconv.Itoa(int(reply.ID)),
			Node:   &reply,
		})
	}

	pageInfo := &model.PageInfo{
		EndCursor:   endCursor,
		HasNextPage: hasNextPage,
	}

	commentConnection := &model.CommentConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}

	return commentConnection, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, commentsDisabled bool) (*model.Post, error) {
	log.Println("mutation resolver: createpost")
	post := model.PostForCreating{
		Title:            title,
		Content:          content,
		CommentsDisabled: commentsDisabled,
	}
	id, err := r.serv.CreatePost(post)
	if err != nil {
		return nil, err
	}
	postToReturn := model.BuildPost(id, post)
	return &postToReturn, nil
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, postID uint, content string, parentID *uint) (*model.Comment, error) {
	log.Println("mutation resolver: create comment")
	var parentIDValue uint
	if parentID == nil {
		parentIDValue = 0
	} else {
		parentIDValue = *parentID
	}
	comment := model.CommentForCreating{
		PostID:    postID,
		Content:   content,
		ParentID:  parentIDValue,
		HasParent: parentID != nil,
	}
	id, err := r.serv.CreateComment(comment)
	if err != nil {
		return nil, err
	}
	commentToReturn := model.BuildComment(id, comment)
	return &commentToReturn, nil
}

// MakeCommentsDisabled is the resolver for the makeCommentsDisabled field.
func (r *mutationResolver) MakeCommentsDisabled(ctx context.Context, postID uint, commentsDisabled bool) (bool, error) {
	log.Println("mutation resolver: makecommentsdis")
	err := r.serv.SetCommentsStatus(postID, commentsDisabled)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Comments is the resolver for the comments field.
func (r *postResolver) Comments(ctx context.Context, obj *model.Post, first *int, after *string) (*model.CommentConnection, error) {
	limit := defaultLimit
	if first != nil {
		limit = *first
	}

	comments, endCursor, hasNextPage, err := r.serv.GetPostComments(obj.ID, limit, after)
	if err != nil {
		return nil, err
	}

	var edges []model.CommentEdge
	for _, comment := range comments {
		edges = append(edges, model.CommentEdge{
			Cursor: strconv.Itoa(int(comment.ID)),
			Node:   &comment,
		})
	}

	pageInfo := &model.PageInfo{
		EndCursor:   endCursor,
		HasNextPage: hasNextPage,
	}

	commentConnection := &model.CommentConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}

	return commentConnection, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, first *int, after *string) (*model.PostConnection, error) {
	log.Println("query resolver: posts")
	limit := defaultLimit
	if first != nil {
		limit = *first
	}

	posts, endCursor, hasNextPage, err := r.serv.GetPosts(limit, after)
	if err != nil {
		return nil, err
	}

	var edges []model.PostEdge
	for _, post := range posts {
		edges = append(edges, model.PostEdge{
			Cursor: strconv.Itoa(int(post.ID)), // Use ID as the cursor
			Node:   &post,
		})
	}

	// Create PageInfo
	pageInfo := &model.PageInfo{
		EndCursor:   endCursor,
		HasNextPage: hasNextPage,
	}

	// Create PostConnection
	postConnection := &model.PostConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}

	return postConnection, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id uint) (*model.Post, error) {
	log.Println("query resolver: post")
	post, err := r.serv.GetPost(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// CommentAdded is the resolver for the commentAdded field.
func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID uint) (<-chan *model.Comment, error) {
	log.Println("sub resolver: commentadded")
	ch := make(chan *model.Comment)
	model.Subs[postID] = append(model.Subs[postID], ch)

	go func() {
		<-ctx.Done()
		close(ch)
	}()
	return ch, nil
}

// Comment returns graph.CommentResolver implementation.
func (r *Resolver) Comment() graph.CommentResolver { return &commentResolver{r} }

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Post returns graph.PostResolver implementation.
func (r *Resolver) Post() graph.PostResolver { return &postResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// Subscription returns graph.SubscriptionResolver implementation.
func (r *Resolver) Subscription() graph.SubscriptionResolver { return &subscriptionResolver{r} }

type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
