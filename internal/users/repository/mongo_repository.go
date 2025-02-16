package repository

import (
	"context"
	"time"

	"github.com/openuniland/good-guy/configs"
	"github.com/openuniland/good-guy/internal/models"
	"github.com/openuniland/good-guy/internal/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName = "test_users"

type userRepo struct {
	cfg         *configs.Configs
	mongoClient *mongo.Client
}

func NewUserRepository(cfg *configs.Configs, mongoClient *mongo.Client) users.Repository {
	return &userRepo{cfg: cfg, mongoClient: mongoClient}
}

func (u *userRepo) Create(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	dbName := u.cfg.MongoDB.MongoDBName

	coll := u.mongoClient.Database(dbName).Collection(collectionName)

	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userRepo) GetUsers(ctx context.Context) ([]*models.User, error) {
	dbName := u.cfg.MongoDB.MongoDBName

	coll := u.mongoClient.Database(dbName).Collection(collectionName)

	filter := bson.D{{
		Key: "is_deleted", Value: false,
	}}
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var users []*models.User
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepo) Find(ctx context.Context, filter interface{}) ([]*models.User, error) {
	dbName := u.cfg.MongoDB.MongoDBName

	coll := u.mongoClient.Database(dbName).Collection(collectionName)

	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var users []*models.User
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
