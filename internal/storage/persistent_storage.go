package storage

import (
	"context"
	"fmt"

	"github.com/unheilbar/artforintrovert_entry_task/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (ps *persistentStorage) Update(ID string, user entities.User) error {
	filter := prepareFilter(ID)
	update := prepareUpdate(user)
	_, err := ps.collection.UpdateOne(context.Background(), filter, update)

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
