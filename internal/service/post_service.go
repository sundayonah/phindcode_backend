package service

import (
	"context"

	"github.com/sundayonah/phindcode_backend/ent"
	"github.com/sundayonah/phindcode_backend/ent/post"
)

type PostService interface {
	CreatePost(ctx context.Context, input *ent.Post) (*ent.Post, error)
	GetPosts(ctx context.Context) ([]*ent.Post, error)
	GetPost(ctx context.Context, id int) (*ent.Post, error)
	UpdatePost(ctx context.Context, id int, input *ent.Post) (*ent.Post, error)
	DeletePost(ctx context.Context, id int) error
}

type postService struct {
	client *ent.Client
}

func NewPostService(client *ent.Client) PostService {
	return &postService{client: client}
}

func (s *postService) CreatePost(ctx context.Context, input *ent.Post) (*ent.Post, error) {
	return s.client.Post.
		Create().
		SetDescription(input.Description).
		SetCategory(input.Category).
		SetCode(input.Code).
		SetImage(input.Image).
		SetUserID(input.UserID).
		Save(ctx)
}

func (s *postService) GetPosts(ctx context.Context) ([]*ent.Post, error) {
	return s.client.Post.
		Query().
		Order(ent.Desc("created_at")).
		All(ctx)
}

func (s *postService) GetPost(ctx context.Context, id int) (*ent.Post, error) {
	return s.client.Post.
		Query().
		Where(post.ID(id)).
		Only(ctx)
}

func (s *postService) UpdatePost(ctx context.Context, id int, input *ent.Post) (*ent.Post, error) {
	return s.client.Post.
		UpdateOneID(id).
		SetDescription(input.Description).
		SetCategory(input.Category).
		SetCode(input.Code).
		SetImage(input.Image).
		Save(ctx)
}

func (s *postService) DeletePost(ctx context.Context, id int) error {
	return s.client.Post.
		DeleteOneID(id).
		Exec(ctx)
}
