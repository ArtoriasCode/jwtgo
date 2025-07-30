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
		Id:    mongoUser.Id.Hex(),
		Email: mongoUser.Email,
		Role:  mongoUser.Role,
		Security: domainEntity.Security{
			Password:     mongoUser.Security.Password,
			Salt:         mongoUser.Security.Salt,
			RefreshToken: mongoUser.Security.RefreshToken,
		},
		CreatedAt: mongoUser.CreatedAt,
		UpdatedAt: mongoUser.UpdatedAt,
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
		Id:    objID,
		Email: domainUser.Email,
		Role:  domainUser.Role,
		Security: mongoEntity.Security{
			Password:     domainUser.Security.Password,
			Salt:         domainUser.Security.Salt,
			RefreshToken: domainUser.Security.RefreshToken,
		},
		CreatedAt: domainUser.CreatedAt,
		UpdatedAt: domainUser.UpdatedAt,
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
	updateFields["role"] = domainUser.Role
	updateFields["security"] = bson.M{
		"password":      domainUser.Security.Password,
		"salt":          domainUser.Security.Salt,
		"refresh_token": domainUser.Security.RefreshToken,
	}
	updateFields["updated_at"] = domainUser.UpdatedAt

	return updateFields
}
