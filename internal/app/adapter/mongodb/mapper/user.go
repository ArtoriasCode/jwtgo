package mapper

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoEntity "jwtgo/internal/app/adapter/mongodb/entity"
	domainEntity "jwtgo/internal/app/entity"
	customErr "jwtgo/internal/error"
)

func MapMongoUserToDomainUser(mongoUser *mongoEntity.User) *domainEntity.User {
	return &domainEntity.User{
		Id:           mongoUser.Id.Hex(),
		Email:        mongoUser.Email,
		Password:     mongoUser.Password,
		Salt:         mongoUser.Salt,
		RefreshToken: mongoUser.RefreshToken,
		CreatedAt:    mongoUser.CreatedAt,
		UpdatedAt:    mongoUser.UpdatedAt,
	}
}

func MapDomainUserToMongoUser(domainUser *domainEntity.User) (*mongoEntity.User, error) {
	var objID primitive.ObjectID
	var err error

	if domainUser.Id != "" {
		objID, err = primitive.ObjectIDFromHex(domainUser.Id)
		if err != nil {
			return nil, customErr.NewInternalServerError("Invalid user ID format")
		}
	} else {
		objID = primitive.NewObjectID()
	}

	return &mongoEntity.User{
		Id:           objID,
		Email:        domainUser.Email,
		Password:     domainUser.Password,
		Salt:         domainUser.Salt,
		RefreshToken: domainUser.RefreshToken,
		CreatedAt:    domainUser.CreatedAt,
		UpdatedAt:    domainUser.UpdatedAt,
	}, nil
}

func MapMongoUsersToDomainUsers(mongoUsers []*mongoEntity.User) []*domainEntity.User {
	var domainUsers []*domainEntity.User
	for _, mongoUser := range mongoUsers {
		domainUsers = append(domainUsers, MapMongoUserToDomainUser(mongoUser))
	}
	return domainUsers
}

func MapDomainUserToBsonUser(domainUser *domainEntity.User) bson.M {
	updateFields := bson.M{}

	if domainUser.Email != "" {
		updateFields["email"] = domainUser.Email
	}
	if domainUser.Password != "" {
		updateFields["password"] = domainUser.Password
	}
	if domainUser.Salt != "" {
		updateFields["salt"] = domainUser.Salt
	}
	if domainUser.RefreshToken != "" {
		updateFields["refresh_token"] = domainUser.RefreshToken
	}
	if !domainUser.UpdatedAt.IsZero() {
		updateFields["updated_at"] = domainUser.UpdatedAt
	}

	return updateFields
}
