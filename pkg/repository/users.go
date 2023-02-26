package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempt" json:"id"`
	Login            string             `bson:"login" json:"login"`
	Password         string             `bson:"password" json:"password"`
	Email            string             `bson:"email" json:"email"`
	IsTemporary      bool               `bson:"is_temporary" json:"is_temporary"`
	Role             int32              `bson:"role" json:"role"`
	RegistrationDate primitive.DateTime `bson:"reg_date" json:"reg_date"`
}
