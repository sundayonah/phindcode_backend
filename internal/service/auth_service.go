package service

import (
	"context"
	"errors"
	"log"

	"github.com/sundayonah/phindcode_backend/ent"
	"github.com/sundayonah/phindcode_backend/ent/user"
)

// AuthService handles user authentication
type AuthService interface {
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
	CreateUser(ctx context.Context, email, name, password string) (*ent.User, error)
	FetchAllUsers(ctx context.Context) ([]*ent.User, error)
}

type authService struct {
	client *ent.Client
}

func NewAuthService(client *ent.Client) AuthService {
	return &authService{client: client}
}

// GetUserByEmail fetches a user by email
func (s *authService) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	user, err := s.client.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

// CreateUser creates a new user
func (s *authService) CreateUser(ctx context.Context, email, fullName, password string) (*ent.User, error) {
	// You can choose to hash the password here if it's provided, or set it to nil for Google login
	user, err := s.client.User.Create().
		SetEmail(email).
		SetFullName(fullName).
		SetPassword(password). // Optional for Google login
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FetchAllUsers fetches all users from the database
func (s *authService) FetchAllUsers(ctx context.Context) ([]*ent.User, error) {
	// Query all users
	users, err := s.client.User.Query().All(ctx)
	if err != nil {
		log.Printf("Failed querying users: %v", err)
		return nil, err
	}

	// Debugging: Print users to the console (optional)
	for _, user := range users {
		log.Printf("ID: %d, Email: %s, Name: %s", user.ID, user.Email, user.FullName)
	}

	return users, nil
}
