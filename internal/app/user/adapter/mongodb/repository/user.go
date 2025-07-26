package repository

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	mongoEntity "jwtgo/internal/app/user/adapter/mongodb/entity"
	"jwtgo/internal/app/user/adapter/mongodb/mapper"
	domainEntity "jwtgo/internal/app/user/entity"
	customErr "jwtgo/internal/pkg/error/type"
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

func (ur *UserRepository) GetById(ctx context.Context, id string) (*domainEntity.User, customErr.BaseErrorInterface) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ur.logger.Error("[UserRepository -> GetById -> ObjectIDFromHex]: ", err)
		return nil, customErr.NewInternalServerError("Failed to get user by id")
	}

	var user mongoEntity.User
	err = ur.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, customErr.NewNotFoundError("User not found")
		}
		ur.logger.Error("[UserRepository -> GetById -> FindOne]: ", err)
		return nil, customErr.NewInternalServerError("Failed to get user by id")
	}

	return mapper.MapMongoUserToDomainUser(&user), nil
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (*domainEntity.User, customErr.BaseErrorInterface) {
	var user mongoEntity.User
	err := ur.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, customErr.NewNotFoundError("User not found")
		}
		ur.logger.Error("[UserRepository -> GetByEmail -> FindOne]: ", err)
		return nil, customErr.NewInternalServerError("Failed to get user by email")
	}

	return mapper.MapMongoUserToDomainUser(&user), nil
}

func (ur *UserRepository) GetAll(ctx context.Context) ([]*domainEntity.User, customErr.BaseErrorInterface) {
	cursor, err := ur.collection.Find(ctx, bson.M{})
	if err != nil {
		ur.logger.Error("[UserRepository -> GetAll -> Find]: ", err)
		return nil, customErr.NewInternalServerError("Failed to get users")
	}

	defer func() {
		if err := cursor.Close(ctx); err != nil {
			ur.logger.Error("[UserRepository -> GetAll -> Close]: ", err)
		}
	}()

	var users []*mongoEntity.User
	if err := cursor.All(ctx, &users); err != nil {
		ur.logger.Error("[UserRepository -> GetAll -> All]: ", err)
		return nil, customErr.NewInternalServerError("Failed to get users")
	}

	return mapper.MapMongoUsersToDomainUsers(users), nil
}

func (ur *UserRepository) Create(ctx context.Context, domainUser *domainEntity.User) (*domainEntity.User, customErr.BaseErrorInterface) {
	mongoUser, err := mapper.MapDomainUserToMongoUser(domainUser)
	if err != nil {
		ur.logger.Error("[UserRepository -> Create -> MapDomainUserToMongoUser]: ", err)
		return nil, customErr.NewInternalServerError("Failed to create user")
	}

	now := time.Now().Unix()
	mongoUser.CreatedAt = now
	mongoUser.UpdatedAt = now

	result, err := ur.collection.InsertOne(ctx, mongoUser)
	if err != nil {
		ur.logger.Error("[UserRepository -> Create -> InsertOne]: ", err)
		return nil, customErr.NewInternalServerError("Failed to create user")
	}

	objID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		ur.logger.Error("[UserRepository -> Create -> InsertedID]: ", "Failed to convert inserted ID to ObjectID")
		return nil, customErr.NewInternalServerError("Failed to create user")
	}

	var createdMongoUser mongoEntity.User
	err = ur.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&createdMongoUser)
	if err != nil {
		ur.logger.Error("[UserRepository -> Create -> FindOne]: ", err)
		return nil, customErr.NewInternalServerError("Failed to create user")
	}

	return mapper.MapMongoUserToDomainUser(&createdMongoUser), nil
}

func (ur *UserRepository) Update(ctx context.Context, id string, domainUser *domainEntity.User) (*domainEntity.User, customErr.BaseErrorInterface) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ur.logger.Error("[UserRepository -> Update -> ObjectIDFromHex]: ", err)
		return nil, customErr.NewInternalServerError("Failed to update user")
	}

	domainUser.UpdatedAt = time.Now().Unix()
	bsonUser := mapper.MapDomainUserToBsonUser(domainUser)

	result, err := ur.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bsonUser})
	if err != nil {
		ur.logger.Error("[UserRepository -> Update -> UpdateOne]: ", err)
		return nil, customErr.NewInternalServerError("Failed to update user")
	}

	if result.MatchedCount == 0 {
		return nil, customErr.NewNotFoundError("User not found")
	}

	var updatedMongoUser mongoEntity.User
	err = ur.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&updatedMongoUser)
	if err != nil {
		ur.logger.Error("[UserRepository -> Update -> FindOne]: ", err)
		return nil, customErr.NewInternalServerError("Failed to update user")
	}

	return mapper.MapMongoUserToDomainUser(&updatedMongoUser), nil
}

func (ur *UserRepository) Delete(ctx context.Context, id string) (bool, customErr.BaseErrorInterface) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ur.logger.Error("[UserRepository -> Delete -> ObjectIDFromHex]: ", err)
		return false, customErr.NewInternalServerError("Failed to delete user")
	}

	result, err := ur.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		ur.logger.Error("[UserRepository -> Delete -> DeleteOne]: ", err)
		return false, customErr.NewInternalServerError("Failed to delete user")
	}

	if result.DeletedCount == 0 {
		return false, customErr.NewNotFoundError("User not found")
	}

	return true, nil
}
