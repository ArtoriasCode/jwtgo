package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	mongoEntity "jwtgo/internal/app/adapter/mongodb/entity"
	"jwtgo/internal/app/adapter/mongodb/mapper"
	domainEntity "jwtgo/internal/app/entity"
	customErr "jwtgo/internal/app/error"
	"jwtgo/pkg/logging"
)

type UserRepository struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func NewUserRepository(client *mongo.Client, database, collection string, logger *logging.Logger) *UserRepository {
	return &UserRepository{
		collection: client.Database(database).Collection(collection),
		logger:     logger,
	}
}

func (ur *UserRepository) GetById(ctx context.Context, id string) (*domainEntity.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, customErr.NewInternalServerError("Invalid user ID format")
	}

	var user mongoEntity.User
	err = ur.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, customErr.NewInternalServerError("Failed to get user")
	}

	return mapper.MapMongoUserToDomainUser(&user), nil
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (*domainEntity.User, error) {
	var user mongoEntity.User
	err := ur.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, customErr.NewInternalServerError("Failed to get user")
	}

	return mapper.MapMongoUserToDomainUser(&user), nil
}

func (ur *UserRepository) GetAll(ctx context.Context) ([]*domainEntity.User, error) {
	cursor, err := ur.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, customErr.NewInternalServerError("Failed to get users")
	}

	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Println("Error closing cursor:", err)
		}
	}()

	var users []*mongoEntity.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return mapper.MapMongoUsersToDomainUsers(users), nil
}

func (ur *UserRepository) Create(ctx context.Context, domainUser *domainEntity.User) (bool, error) {
	mongoUser, err := mapper.MapDomainUserToMongoUser(domainUser)
	if err != nil {
		ur.logger.Error("Error while mapping user: ", err)
		return false, err
	}

	_, err = ur.collection.InsertOne(ctx, mongoUser)
	if err != nil {
		return false, customErr.NewInternalServerError("Failed to create a user")
	}

	return true, nil
}

func (ur *UserRepository) Update(ctx context.Context, id string, domainUser *domainEntity.User) (bool, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, customErr.NewInternalServerError("Invalid user ID format")
	}

	domainUser.UpdatedAt = time.Now().UTC()
	bsonUser := mapper.MapDomainUserToBsonUser(domainUser)

	_, err = ur.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bsonUser})
	if err != nil {
		return false, customErr.NewInternalServerError("Failed to update user")
	}

	return true, nil
}

func (ur *UserRepository) Delete(ctx context.Context, id string) (bool, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, customErr.NewInternalServerError("Invalid user ID format")
	}

	_, err = ur.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return false, customErr.NewInternalServerError("Failed to delete user")
	}

	return true, nil
}
