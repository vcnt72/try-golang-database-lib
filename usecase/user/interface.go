package user

import (
	"context"

	"github.com/vcnt72/try-golang-database-lib/entity"
)

type Repository interface {
	Store(ctx context.Context, user *entity.User) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByID(ctx context.Context, id string) (*entity.User, error)
}

type Usecase interface {
	Register(ctx context.Context, createDTO CreateUserDTO) (*entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Login(ctx context.Context, email, password string) (user *entity.User, tokenStr string, err error)
}
