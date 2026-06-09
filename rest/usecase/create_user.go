package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sharma-ayush1999/rest/domain"
)

type CreateUserInput struct {
	Email 	string
	Name 	string
}

type CreateUserOutput struct {
	UserID string
}

type CreateUserUseCase struct {
	repo domain.UserRepository	// depends on abstraction, not implementation
	idGen func() string	 		// injected ID generator (testable)
}

func NewCreateUserUserCase(repo domain.UserRepository, idGen func() string) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo, idGen: idGen}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
	user := &domain.User {
		ID: uc.idGen(),
		Email: input.Email,
		Name: input.Name,
		CreatedAt: time.Now(),
	}

	// Domain validation
	if err := user.Validate(); err != nil {
		return nil, fmt.Errorf("Validation failed: %w", err)
	}

	// Check duplicate
	existing, err := uc.repo.FindByEmail(ctx, input.Email)
	if err != nil && existing != nil {
		return nil, errors.New("email already registered")
	}

	if err := uc.repo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	
	return &CreateUserOutput{UserID: user.ID}, nil
}


