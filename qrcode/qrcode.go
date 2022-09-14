package qrcode

import (
	"context"
	"os"
	"server/database"
	"server/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func loadUsers(qrcode string) (*model.UserSmall, error) {
	_collections_users := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("users")
	var _user *model.UserSmall
	err := _collections_users.FindOne(context.Background(), bson.D{{Key: "keypair.publicKey", Value: qrcode}, {Key: "deleted", Value: false}}).Decode(&_user)
	if err != nil {
		return nil, err
	}

	return _user, nil
}

func loadTransaction(qrcode string) (*model.Paiement, error) {
	if primitive.IsValidObjectID(qrcode) {
		objectId, _ := primitive.ObjectIDFromHex(qrcode)
		_collections_transactions := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("transactions")
		var _transaction *model.Paiement
		err := _collections_transactions.FindOne(context.Background(), bson.D{{Key: "_id", Value: objectId}, {Key: "status", Value: bson.D{{Key: "$in", Value: bson.A{model.PaymentStatusRequiresAction.String(), model.PaymentStatusRequiresConfirmation.String(), model.PaymentStatusRequiresPaiement.String()}}}}}).Decode(&_transaction)
		if err != nil {
			return nil, err
		}
		return _transaction, nil
	}

	return nil, nil

}

func loadEnterprise(qrcode string) (*model.EnterpriseSmall, error) {
	_collections_enterprises := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("entreprises")
	var _enterpriseSmall *model.EnterpriseSmall
	err := _collections_enterprises.FindOne(context.Background(), bson.D{{Key: "walletPublicKey", Value: qrcode}, {Key: "removed", Value: false}}).Decode(&_enterpriseSmall)
	if err != nil {
		return nil, err
	}
	return _enterpriseSmall, nil
}

func QueryQrCodeUsers(qrcode string) (model.QRCodeOwner, error) {
	user, _ := loadUsers(qrcode)
	enterprise, _ := loadEnterprise(qrcode)
	transaction, _ := loadTransaction(qrcode)
	if user != nil {
		return user, nil
	}

	if enterprise != nil {
		return enterprise, nil
	}

	if transaction != nil {
		return transaction, nil
	}

	return nil, nil

}
