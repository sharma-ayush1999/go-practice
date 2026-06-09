package domain

import (
	"context"
	"errors"
	"regexp"
	"time"
)


type User struct {
	ID			string
	Email		string
	Name		string
	CreatedAt	time.Time
}

// Domain validation — business rules live here, not in handlers
func (u *User) Validate() error {
	if u.Email == " " {
		return errors.New("email is required")
	}
	emailRe := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,}$`)
	if !emailRe.MatchString(u.Email){
		return errors.New("invalid email format")
	}
	if len(u.Name) < 2 {
		return errors.New("name must be atleast 2 characters")
	}

	return nil
}

// UserRepository — interface defined in domain, implemented in infra layer
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
}