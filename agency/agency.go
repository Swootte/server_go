package agency

import (
	"context"
	"log"
	"os"
	"server/database"
	"server/errors"
	"server/graph/model"
	"server/utils"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	utils.LoadEnv()
}

func RetrieveAllAgencies(ctx context.Context) ([]*model.Agency, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("agencies")
	cursor, err := _collections.Aggregate(ctx, mongo.Pipeline{bson.D{{Key: "$limit", Value: 30}}})
	if err != nil {
		return nil, err
	}

	var agencies []*model.Agency
	for cursor.Next(ctx) {
		var singleAgency *model.Agency
		if err = cursor.Decode(&singleAgency); err != nil {
			errors.SaveError(ctx, err, "34", "agency", "", "")
			log.Fatal(err)
		}
		agencies = append(agencies, singleAgency)
	}

	return agencies, nil

}

func CreateAgency(ctx context.Context, agency model.AgencyInpyt, creator string, ip string) (string, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("agencies")
	objectId, _ := primitive.ObjectIDFromHex(creator)
	_time := time.Now().UTC().Format(time.RFC3339)
	_agency := database.DBAgency{
		ShortId:   uuid.NewString(),
		Title:     agency.Title,
		Address:   agency.Address,
		Status:    model.AgencyOpenStatusOpen.String(),
		Creator:   objectId,
		CreatedAt: _time,
		UpdatedAt: _time,
		City:      agency.City,
		Country:   agency.Country,
	}

	result, err := _collections.InsertOne(ctx, _agency)
	if err != nil {
		errors.SaveError(ctx, err, "61", "agency", "", "")
		return "", err
	}

	insertedID := result.InsertedID.(primitive.ObjectID).Hex()

	return insertedID, nil
}
