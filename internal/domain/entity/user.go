package entity

import (
	"time"
)

type User struct {
	Id           string    `bson:"_id,omitempty" json:"id"`
	Email        string    `bson:"email" json:"email"`
	Password     string    `bson:"password" json:"password"`
	Salt         string    `bson:"salt" json:"salt"`
	RefreshToken string    `bson:"refresh_token" json:"refresh_token"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at" json:"updated_at"`
}
