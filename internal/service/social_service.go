package service

import (
	"context"

	"github.com/sundayonah/phindcode_backend/ent"
	"github.com/sundayonah/phindcode_backend/ent/post"
	"github.com/sundayonah/phindcode_backend/ent/like"
	"github.com/sundayonah/phindcode_backend/ent/comment"
	"github.com/sundayonah/phindcode_backend/ent/share"
)

type SocialService interface {
	// Like operations
	LikePost(ctx context.Context, postID int, userID string) error
	UnlikePost(ctx context.Context, postID int, userID string) error
	GetPostLikes(ctx context.Context, postID int) ([]*ent.Like, error)

	// Comment operations
	CreateComment(ctx context.Context, postID int, userID string, content string) (*ent.Comment, error)
	UpdateComment(ctx context.Context, commentID int, content string) (*ent.Comment, error)
	DeleteComment(ctx context.Context, commentID int) error
	GetPostComments(ctx context.Context, postID int) ([]*ent.Comment, error)

	// Share operations
	SharePost(ctx context.Context, postID int, userID string, shareTo string) (*ent.Share, error)
	GetPostShares(ctx context.Context, postID int) ([]*ent.Share, error)
}

type socialService struct {
	client *ent.Client
}

func NewSocialService(client *ent.Client) SocialService {
	return &socialService{client: client}
}

// Like implementations
func (s *socialService) LikePost(ctx context.Context, postID int, userID string) error {
	return s.client.Like.Create().
		SetUserID(userID).
		SetPostID(postID).
		Exec(ctx)
}

func (s *socialService) UnlikePost(ctx context.Context, postID int, userID string) error {
	_, err := s.client.Like.Delete().
		Where(
			like.And(
				like.HasPostWith(post.ID(postID)),
				like.UserID(userID),
			),
		).
		Exec(ctx)
	return err
}

func (s *socialService) GetPostLikes(ctx context.Context, postID int) ([]*ent.Like, error) {
	return s.client.Like.Query().
		Where(like.HasPostWith(post.ID(postID))).
		All(ctx)
}

// Comment implementations
func (s *socialService) CreateComment(ctx context.Context, postID int, userID string, content string) (*ent.Comment, error) {
	return s.client.Comment.Create().
		SetContent(content).
		SetUserID(userID).
		SetPostID(postID).
		Save(ctx)
}

func (s *socialService) UpdateComment(ctx context.Context, commentID int, content string) (*ent.Comment, error) {
	return s.client.Comment.UpdateOneID(commentID).
		SetContent(content).
		Save(ctx)
}

func (s *socialService) DeleteComment(ctx context.Context, commentID int) error {
	return s.client.Comment.DeleteOneID(commentID).Exec(ctx)
}

func (s *socialService) GetPostComments(ctx context.Context, postID int) ([]*ent.Comment, error) {
	return s.client.Comment.Query().
		Where(comment.HasPostWith(post.ID(postID))).
		Order(ent.Desc(comment.FieldCreatedAt)).
		All(ctx)
}

// Share implementations
func (s *socialService) SharePost(ctx context.Context, postID int, userID string, shareTo string) (*ent.Share, error) {
	return s.client.Share.Create().
		SetUserID(userID).
		SetShareTo(shareTo).
		SetPostID(postID).
		Save(ctx)
}

func (s *socialService) GetPostShares(ctx context.Context, postID int) ([]*ent.Share, error) {
	return s.client.Share.Query().
		Where(share.HasPostWith(post.ID(postID))).
		All(ctx)
}
