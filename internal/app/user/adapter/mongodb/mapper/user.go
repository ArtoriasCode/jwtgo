package mapper

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	mongoEntity "jwtgo/internal/app/user/adapter/mongodb/entity"
	domainEntity "jwtgo/internal/app/user/entity"
	customErr "jwtgo/internal/pkg/error/type"
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

	updateFields["email"] = domainUser.Email
	updateFields["password"] = domainUser.Password
	updateFields["salt"] = domainUser.Salt
	updateFields["refresh_token"] = domainUser.RefreshToken
	updateFields["updated_at"] = domainUser.UpdatedAt

	return updateFields
}
