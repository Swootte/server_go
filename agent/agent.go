package agent

import (
	"context"
	"fmt"
	"os"
	"server/database"
	"server/errors"
	"server/finance"
	"server/graph/model"
	_user "server/user"
	"server/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	utils.LoadEnv()
}

func AssignRoleToAddress(ctx context.Context, account string, role string, user model.User, ip string) (*bool, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	user_to_Receive, _ := _user.GetUserById(account)
	_, err := finance.GiveRole(user_to_Receive.Keypair.PublicKey, role, user.Keypair.SecretKey, user.Keypair.PublicKey)
	if err != nil {
		errors.SaveError(ctx, err, "27", "agent", "", "")
		return nil, err
	}

	objectId, _ := primitive.ObjectIDFromHex(account)
	_time := time.Now().UTC().Format(time.RFC3339)

	update := bson.M{
		"$set":  bson.D{{Key: "updatedAt", Value: _time}},
		"$push": bson.D{{Key: "permissions", Value: role}},
	}

	res, err := _collections.UpdateOne(ctx, bson.D{{Key: "_id", Value: objectId}}, update)
	if err != nil {
		errors.SaveError(ctx, err, "42", "agent", "", "")
		return nil, err
	}

	if res.ModifiedCount == 0 {
		errors.SaveError(ctx, err, "47", "agent", "", "")
		return nil, fmt.Errorf("no error was defined")
	}

	result := true

	return &result, nil

}

func UnAssignRoleToAddress(ctx context.Context, account string, role string, user model.User, ip string) (*bool, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	user_to_Receive, _ := _user.GetUserById(account)
	_, err := finance.RemoveRole(user_to_Receive.Keypair.PublicKey, role, user.Keypair.SecretKey, user.Keypair.PublicKey)
	if err != nil {
		return nil, err
	}

	objectId, _ := primitive.ObjectIDFromHex(account)
	_time := time.Now().UTC().Format(time.RFC3339)

	update := bson.M{
		"$set":  bson.D{{Key: "updatedAt", Value: _time}},
		"$pull": bson.D{{Key: "permissions", Value: role}},
	}
	res, err := _collections.UpdateOne(ctx, bson.D{{Key: "_id", Value: objectId}}, update)
	if err != nil {
		errors.SaveError(ctx, err, "74", "agent", "", "")
		return nil, err
	}

	if res.ModifiedCount == 0 {
		return nil, fmt.Errorf("no document was found")
	}

	result := true

	return &result, nil
}

func CancelTransaction(ctx context.Context, transactionID string, typeArg model.PaymentType, pinCode string, canceller model.User, ip string) (bool, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(transactionID)
	cancellerObjectId, _ := primitive.ObjectIDFromHex(canceller.ID)
	_time := time.Now().UTC().Format(time.RFC3339)
	if typeArg == model.PaymentTypeTopup {
		_, err := finance.CancelWithDraw(transactionID, canceller.Keypair.SecretKey, canceller.Keypair.PublicKey)
		if err != nil {
			errors.SaveError(ctx, err, "95", "agent", "", "")
			return false, err
		}
	} else if typeArg == model.PaymentTypeWithdraw {
		_, err := finance.CancelDeposit(transactionID, canceller.Keypair.SecretKey, canceller.Keypair.PublicKey)
		if err != nil {
			errors.SaveError(ctx, err, "101", "agent", "", "")
			return false, err
		}
	}

	_, err := _collections.UpdateOne(ctx, bson.D{{Key: "_id", Value: objectId}}, bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: model.PaymentStatusCancelledAgent}, {Key: "updatedAt", Value: _time}, {Key: "cancellor", Value: cancellerObjectId}}}})
	if err != nil {
		errors.SaveError(ctx, err, "108", "agent", "", "")
		return false, err
	}
	return true, nil
}

func ConfirmTransaction(ctx context.Context, transactionID string, typeArg model.PaymentType, pinCode string, approver model.User, ip string) (bool, error) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
	objectId, _ := primitive.ObjectIDFromHex(transactionID)
	approverObjectId, _ := primitive.ObjectIDFromHex(approver.ID)
	_time := time.Now().UTC().Format(time.RFC3339)
	if typeArg == model.PaymentTypeTopup {
		_, err := finance.ApproveDeposit(transactionID, approver.Keypair.SecretKey, approver.Keypair.PublicKey)
		if err != nil {
			errors.SaveError(ctx, err, "122", "agent", "", "")
			return false, err
		}
	} else if typeArg == model.PaymentTypeWithdraw {
		_, err := finance.ApproveWithdraw(transactionID, approver.Keypair.SecretKey, approver.Keypair.PublicKey)
		if err != nil {
			errors.SaveError(ctx, err, "128", "agent", "", "")
			return false, err
		}
	}

	_, err := _collections.UpdateOne(ctx, bson.D{primitive.E{Key: "_id", Value: objectId}}, bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "status", Value: model.PaymentStatusDone}, primitive.E{Key: "updatedAt", Value: _time}, primitive.E{Key: "validatorId", Value: approverObjectId}}}})
	if err != nil {
		errors.SaveError(ctx, err, "135", "agent", "", "")
		return false, err
	}
	return true, nil

}
