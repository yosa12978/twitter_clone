package repos

import (
	"context"

	"github.com/yosa12978/twitter/user-api/db"
	"github.com/yosa12978/twitter/user-api/logging"
	"github.com/yosa12978/twitter/user-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User interface {
	FindAll(ctx context.Context) ([]types.User, error)
	FindById(ctx context.Context, id string) (*types.User, error)
	FindByEmail(ctx context.Context, email string) (*types.User, error)
	FindByName(ctx context.Context, username string) (*types.User, error)
	FindByCredentials(ctx context.Context, username string, passwordHash string) (*types.User, error)
	FindByCredentialsEmail(ctx context.Context, email string, passwordHash string) (*types.User, error)
	Create(ctx context.Context, user types.User) (string, error)
	Update(ctx context.Context, id string, user types.User) (string, error)
	Remove(ctx context.Context, id string) (string, error)
}

type userMongo struct {
	logger logging.Logger
	db     *mongo.Database
}

func New(ctx context.Context) User {
	repo := new(userMongo)
	repo.logger = logging.New("userRepoMongo")
	repo.db = db.GetDB(ctx)
	return repo
}

func (repo *userMongo) FindAll(ctx context.Context) ([]types.User, error) {
	var users []types.User
	opts := options.Find().SetSort(bson.M{"username": 1})
	cursor, err := repo.db.Collection("users").Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *userMongo) FindById(ctx context.Context, id string) (*types.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user types.User
	err = repo.db.Collection("users").FindOne(ctx, bson.M{"baseModel._id": objId}).Decode(&user)
	return &user, err
}

func (repo *userMongo) FindByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User
	err := repo.db.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return &user, err
}

func (repo *userMongo) FindByName(ctx context.Context, username string) (*types.User, error) {
	var user types.User
	err := repo.db.Collection("users").FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return &user, err
}

func (repo *userMongo) FindByCredentials(ctx context.Context, username string, passwordHash string) (*types.User, error) {
	var user types.User
	q := bson.M{
		"$and": bson.A{
			bson.M{"username": username},
			bson.M{"passwordHash": passwordHash},
		},
	}
	err := repo.db.Collection("users").FindOne(ctx, q).Decode(&user)
	return &user, err
}

func (repo *userMongo) FindByCredentialsEmail(ctx context.Context, email string, passwordHash string) (*types.User, error) {
	var user types.User
	q := bson.M{
		"$and": bson.A{
			bson.M{"email": email},
			bson.M{"passwordHash": passwordHash},
		},
	}
	err := repo.db.Collection("users").FindOne(ctx, q).Decode(&user)
	return &user, err
}

func (repo *userMongo) Create(ctx context.Context, user types.User) (string, error) {
	res, err := repo.db.Collection("users").InsertOne(ctx, user)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (repo *userMongo) Update(ctx context.Context, id string, user types.User) (string, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	_, err = repo.db.Collection("users").UpdateByID(ctx, objId, user)
	return id, err
}

func (repo *userMongo) Remove(ctx context.Context, id string) (string, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	_, err = repo.db.Collection("users").DeleteOne(context.TODO(), bson.M{"_id": objId})
	return id, err
}
