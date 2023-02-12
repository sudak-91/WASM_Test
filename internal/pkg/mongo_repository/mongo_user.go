package mongorepository

import (
	"context"
	"fmt"
	"log"

	"github.com/sudak-91/wasm-test/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUser struct {
	collection *mongo.Collection
}

func NewMongoUser(db *mongo.Database) *MongoUser {
	var (
		mongoUser = &MongoUser{
			collection: db.Collection("User"),
		}
	)
	return mongoUser
}

func (u *MongoUser) CreateUser(user repository.User) error {
	data, err := bson.Marshal(user)
	if err != nil {
		log.Fatalf("Users Marshaling has error: %s", err.Error())
		return err
	}
	result, err := u.collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatalf("InsertOne method has error %s", err.Error())
		return err
	}
	resultID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatal("Convert is not ok")
		return fmt.Errorf("Result ID has error")
	}
	log.Println(resultID)
	return nil

}

func (u *MongoUser) ReadUserByEmail(email string) ([]repository.User, error) {
	filter := bson.D{{"email", email}}
	cur, err := u.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var users []repository.User
	err = cur.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *MongoUser) ReadUser(login string) ([]repository.User, error) {
	filter := bson.D{{"login", login}}
	result := u.collection.FindOne(context.TODO(), filter)
	var User []repository.User
	err := result.Decode(&User)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (u *MongoUser) UpdateUser(newUser repository.User) error {
	filter := bson.D{{"_id", newUser.ID}}
	data, err := bson.Marshal(newUser)
	if err != nil {
		return err
	}
	u.collection.UpdateOne(context.TODO(), filter, data)
	return nil
}

func (u *MongoUser) DeleteUser(user repository.User) error {
	resul, err := u.collection.DeleteOne(context.TODO(), bson.D{{"_id", user.ID}})
	if err != nil {
		return err
	}
	fmt.Printf("resul: %v\n", resul)
	return nil
}
