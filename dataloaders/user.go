package dataloaders

// import graph gophers with your other imports
import (
	"context"
	"fmt"
	"log"
	"os"
	"server/database"
	"server/graph/model"

	"github.com/graph-gophers/dataloader"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	usersIds := bson.A{}
	for _, key := range keys {
		objectId, _ := primitive.ObjectIDFromHex(key.String())
		usersIds = append(usersIds, objectId)
	}

	_match := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: usersIds}}}}
	cursor, err := _collections.Find(ctx, _match)
	if err != nil {
		return nil
	}

	userById := map[string]*model.UserSmall{}
	for cursor.Next(ctx) {
		var singleUser *model.UserSmall
		if err = cursor.Decode(&singleUser); err != nil {
			log.Fatal(err)
		}
		userById[singleUser.ID] = singleUser
	}

	// return users in the same order requested
	output := make([]*dataloader.Result, len(keys))
	for index, userKey := range keys {
		user, ok := userById[userKey.String()]
		if ok {
			output[index] = &dataloader.Result{Data: user, Error: nil}
		} else {
			err := fmt.Errorf("user not found %s", userKey.String())
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output
}

func (l *Loaders) User(ctx context.Context, userID string) (*model.UserSmall, error) {
	thunk := l.UserLoader.Load(ctx, dataloader.StringKey(userID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*model.UserSmall), nil
}
