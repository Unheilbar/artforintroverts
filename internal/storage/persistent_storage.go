package storage

import (
	"context"
	"fmt"

	"github.com/unheilbar/artforintrovert_entry_task/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type persistentStorage struct {
	collection *mongo.Collection
}

func (ps *persistentStorage) GetAll() ([]entities.User, error) {
	cur, err := ps.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var users []entities.User

	err = cur.All(context.Background(), &users)

	return users, err
}

func (ps *persistentStorage) Delete(ID string) error {
	filter := prepareFilter(ID)

	deleted, err := ps.collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	if deleted.DeletedCount == 0 {
		return fmt.Errorf("couldn't delete user with %s", ID)
	}

	return nil
}

func (ps *persistentStorage) UpdateAndReturn(ID string, user entities.User) (entities.User, error) {
	filter := prepareFilter(ID)
	update := prepareUpdate(user)
	after := options.After
	res := ps.collection.FindOneAndUpdate(context.Background(), filter, update,
		&options.FindOneAndUpdateOptions{ReturnDocument: &after})

	if err := res.Err(); err != nil {
		return entities.User{}, res.Err()
	}

	var updated entities.User
	err := res.Decode(&updated)

	return updated, err
}

func (ps *persistentStorage) GetById(ID string) (entities.User, error) {
	filter := prepareFilter(ID)

	res := ps.collection.FindOne(context.Background(), filter)

	if err := res.Err(); err != nil {
		return entities.User{}, res.Err()
	}

	var user entities.User
	err := res.Decode(&user)

	return user, err
}

func (ps *persistentStorage) Save(user entities.User) error {
	buser, err := bson.Marshal(user)
	if err != nil {
		return err
	}

	_, err = ps.collection.InsertOne(context.Background(), buser)

	return err
}

func prepareUpdate(user entities.User) bson.M {
	bUser, _ := bson.Marshal(user)

	var update bson.M

	bson.Unmarshal(bUser, &update)

	return bson.M{
		"$set": update,
	}
}

func prepareFilter(ID string) bson.M {
	return bson.M{
		"id": ID,
	}
}
