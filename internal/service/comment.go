package service

import (
	"context"
	"go-rest-api-app/internal/domain"
)

// Stor interface to service
type Stor interface {
	GetComment(ctx context.Context, id string) (domain.Comment, error)
	AddComment(ctx context.Context, comment domain.Comment) error
	UpdateComment(ctx context.Context, comment domain.Comment) error
	DeleteComment(ctx context.Context, id string) (domain.Comment, error)
}

// Comment service
type ServiceComment struct {
	Stor
}

// constructor
func NewServiceComment(stor Stor) *ServiceComment {
	return &ServiceComment{Stor: stor}
}

// result slice Comment
// func (s *Service) GetComments(ctx context.Context) ([]Comment, error) {
// 	return []Comment{}, nil
// }

// result one comment to id
func (s *ServiceComment) GetComment(ctx context.Context, id string) (domain.Comment, error) {
	cmt, err := s.Stor.GetComment(ctx, id)
	if err != nil {
		return domain.Comment{}, err
	}
	return cmt, nil
}

// Add new Comment
func (s *ServiceComment) AddComment(ctx context.Context, comment domain.Comment) error {
	return nil
}

// Update Comment
func (s *ServiceComment) UpdateComment(ctx context.Context, comment domain.Comment) error {
	return nil
}

// Delete Comment
func (s *ServiceComment) DeleteComment(ctx context.Context, id string) (domain.Comment, error) {
	return domain.Comment{}, nil
}
