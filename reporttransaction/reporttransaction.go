package reporttransaction

import (
	"context"
	"os"
	"server/database"
	"server/graph/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReportTransaction(user model.User, transaction_id string, message string, ip string) (*bool, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("reportTransactions")
	_time := time.Now().UTC().Format(time.RFC3339)
	tx, _ := primitive.ObjectIDFromHex(transaction_id)
	reportedBy, _ := primitive.ObjectIDFromHex(user.ID)

	_report := database.DBReportTransaction{
		Transaction: tx,
		Message:     message,
		Status:      "",
		ReportedBy:  reportedBy,
		RevolvedBy:  nil,
		CreatedAt:   _time,
		UpdatedAt:   _time,
	}
	return_ := true

	_, err := _collections.InsertOne(context.Background(), _report)
	if err != nil {
		return nil, err
	}

	return &return_, nil

}
