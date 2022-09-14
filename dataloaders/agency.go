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

func GetAgencies(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("agencies")
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

	agencies := map[string]*model.Agency{}
	for cursor.Next(ctx) {
		var singleAgency *model.Agency
		if err = cursor.Decode(&singleAgency); err != nil {
			log.Fatal(err)
		}
		agencies[*singleAgency.ID] = singleAgency
	}

	// return users in the same order requested
	output := make([]*dataloader.Result, len(keys))
	for index, agencyKey := range keys {
		user, ok := agencies[agencyKey.String()]
		if ok {
			output[index] = &dataloader.Result{Data: user, Error: nil}
		} else {
			err := fmt.Errorf("user not found %s", agencyKey.String())
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output

}

func (l *Loaders) Agency(ctx context.Context, agencyID string) (*model.Agency, error) {
	thunk := l.AgencyLoader.Load(ctx, dataloader.StringKey(agencyID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*model.Agency), nil
}
