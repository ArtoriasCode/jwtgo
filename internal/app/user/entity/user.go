package entity

type User struct {
	Id           string `bson:"_id,omitempty" json:"id"`
	Email        string `bson:"email" json:"email"`
	Password     string `bson:"password" json:"password"`
	Salt         string `bson:"salt" json:"salt"`
	RefreshToken string `bson:"refresh_token" json:"refresh_token"`
	CreatedAt    int64  `bson:"created_at" json:"created_at"`
	UpdatedAt    int64  `bson:"updated_at" json:"updated_at"`
}
