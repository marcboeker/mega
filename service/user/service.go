package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/marcboeker/mega/db"
	"github.com/marcboeker/mega/ent"
	"github.com/marcboeker/mega/graph/model"
)

// The Service interface defines all available methods the service offers.
type Service interface {
	Get(ctx context.Context, id uuid.UUID) (*ent.User, error)
	Create(ctx context.Context, user model.AddUserInput) (*ent.User, error)
}

// The User service acts as a wrapper for the service methods.
type User struct {
	client *db.Client
}

// New initializes the user service with the given DB client.
func New(client *db.Client) Service {
	return &User{client: client}
}

// Get retrieves a user by its ID.
func (s *User) Get(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return s.client.User.Get(ctx, id)
}

// Create saves a new user with the given input.
func (s *User) Create(ctx context.Context, user model.AddUserInput) (*ent.User, error) {
	return s.client.User.
		Create().
		SetName(user.Name).
		SetAge(user.Age).
		Save(ctx)
}
