package repos

import (
	"context"

	"github.com/yosa12978/twitter/user-api/logging"
	"github.com/yosa12978/twitter/user-api/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	FindAll(ctx context.Context) ([]types.User, error)
	FindById(ctx context.Context, id string) (*types.User, error)
	FindByEmail(ctx context.Context, email string) (*types.User, error)
	FindByName(ctx context.Context, username string) (*types.User, error)
	FindByCredentials(ctx context.Context, username string, password string) (*types.User, error)
	FindByCredentialsEmail(ctx context.Context, email string, password string) (*types.User, error)
	Create(ctx context.Context, user types.User) (string, error)
	Update(ctx context.Context, id string, user types.User) (string, error)
	Remove(ctx context.Context, id string) (string, error)
}

type userRepoMongo struct {
	logger logging.Logger
	db     *mongo.Database
}

func (repo *userRepoMongo) FindAll(ctx context.Context) ([]types.User, error) {
	return nil, nil
}

func (repo *userRepoMongo) FindById(ctx context.Context, id string) (*types.User, error) {
	return nil, nil
}

func (repo *userRepoMongo) FindByEmail(ctx context.Context, email string) (*types.User, error) {
	return nil, nil
}

func (repo *userRepoMongo) FindByName(ctx context.Context, username string) (*types.User, error) {
	return nil, nil
}

func (repo *userRepoMongo) FindByCredentials(ctx context.Context, username string, password string) (*types.User, error) {
	return nil, nil
}

func (repo *userRepoMongo) FindByCredentialsEmail(ctx context.Context, email string, password string) (*types.User, error) {
	return nil, nil
}

func (repo *userRepoMongo) Create(ctx context.Context, user types.User) (string, error) {
	return "", nil
}

func (repo *userRepoMongo) Update(ctx context.Context, id string, user types.User) (string, error) {
	return "", nil
}

func (repo *userRepoMongo) Remove(ctx context.Context, id string) (string, error) {
	return "", nil
}
