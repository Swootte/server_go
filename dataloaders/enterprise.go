package dataloaders

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

func GetEnterprises(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	userIDs := bson.A{}
	for _, key := range keys {
		objectId, _ := primitive.ObjectIDFromHex(key.String())
		userIDs = append(userIDs, objectId)
	}

	_match := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: userIDs}}}}
	cursor, err := _collections.Find(ctx, _match)
	if err != nil {
		return nil
	}

	userById := map[string]*model.EnterpriseSmall{}
	for cursor.Next(ctx) {
		var singleEnterprise *model.EnterpriseSmall
		if err = cursor.Decode(&singleEnterprise); err != nil {
			log.Fatal(err)
		}
		userById[singleEnterprise.ID] = singleEnterprise
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

func (l *Loaders) Enterprise(ctx context.Context, enterpriseID string) (*model.EnterpriseSmall, error) {
	thunk := l.EnterpriseLoader.Load(ctx, dataloader.StringKey(enterpriseID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*model.EnterpriseSmall), nil
}
