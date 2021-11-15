package repository

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/vcnt72/try-golang-database-lib/entity"
	"github.com/vcnt72/try-golang-database-lib/usecase/user"
)

type userInMemRepository struct {
	users []*entity.User
}

func NewUserInMemRepository() user.Repository {
	return &userInMemRepository{}
}

func (m *userInMemRepository) Store(ctx context.Context, user *entity.User) (*entity.User, error) {

	user.ID = uuid.NewV4().String()
	m.users = append(m.users, user)
	return user, nil
}

func (m *userInMemRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {

	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, entity.ErrNotFound
}

func (m *userInMemRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	panic("not implemented") // TODO: Implement
}
