package notification

import (
	"context"
	"log"
	"os"
	"server/database"
	snippets "server/firebase"
	"server/graph/model"
	"server/utils"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	utils.LoadEnv()
}

func LoadNotification(ctx context.Context, userID string) ([]*model.Notification, error) {
	objectId, _ := primitive.ObjectIDFromHex(userID)
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("notifications")
	cursor, err := _collections.Find(ctx, bson.M{"to": objectId})
	if err != nil {
		return nil, err
	}

	var notifications []*model.Notification
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var singleNotification *model.Notification
		if err = cursor.Decode(&singleNotification); err != nil {
			log.Fatal(err)
		}
		notifications = append(notifications, singleNotification)
	}

	return notifications, nil
}

func LoadNotificationCount(ctx context.Context, userID string) (*float64, error) {
	objectId, _ := primitive.ObjectIDFromHex(userID)
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("notifications")
	count, err := _collections.CountDocuments(ctx, bson.M{"to": objectId, "isRead": true})
	if err != nil {
		return nil, err
	}
	result := float64(count)
	return &result, nil

}

func SetAllNotificationToRead(ctx context.Context, userID string) (bool, error) {
	objectId, _ := primitive.ObjectIDFromHex(userID)
	_time := time.Now().UTC().Format(time.RFC3339)
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("notifications")
	_, err := _collections.UpdateMany(ctx, bson.M{"to": objectId, "isRead": false}, bson.M{"isRead": true, "updatedAt": _time})
	if err != nil {
		return false, err
	}
	return true, nil
}

func SaveNotification(ctx context.Context, notification model.Notification, _to string) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("notifications")
	_time := time.Now().UTC().Format(time.RFC3339)
	from, _ := primitive.ObjectIDFromHex(*notification.FromID)
	to, _ := primitive.ObjectIDFromHex(_to)
	_notification := database.DBNotification{
		Text:      *notification.Text,
		Type:      *notification.Type,
		ImgUrl:    *notification.ImgURL,
		IsRead:    false,
		From:      from,
		To:        to,
		CreatedAt: _time,
	}
	_collections.InsertOne(ctx, _notification)
}

func build_and_send_notification(sender string, receiver model.User, transaction model.Paiement, text_database string, text_notif string, _type string) {
	_testing, _ := strconv.ParseBool(os.Getenv("testing"))
	if !_testing {
		snippets.Connect().SendNotificationMEssage(_type, _type, text_notif, transaction.ID, *receiver.FcmToken)
	}
}

func CreateDBNotification(_type model.PaymentType, sender model.User, receiver model.User, transaction model.Paiement) {
	switch _type {
	case model.PaymentTypeTopup:
		build_and_send_notification(sender.ID, receiver, transaction, `Votre dépot d'un montant de ${transaction.amount} à bien été effectué rendez vous dans votre portefeuille pour le constater`, `Votre dépot d'un montant de ${transaction.amount} à bien été effectué rendez vous dans votre portefeuille pour le constater`, transaction.Type.String())
	case model.PaymentTypeCommerce:
		build_and_send_notification(sender.ID, receiver, transaction, `${sender.first_name} ${sender.last_name}  vient de vous faire paiement d'un montant de ${transaction.amount} cfa`, `${sender.first_name}  ${sender.last_name}  vient de vous payer ${transaction.amount} cfa`, transaction.Type.String())
	case model.PaymentTypeWithdraw:
		build_and_send_notification(sender.ID, receiver, transaction, `Votre de retrait d'un montant de ${transaction.amount} à bien été effectué`, `Votre dépot d'un montant de ${transaction.amount} à bien été effectué`, "WITHDRAW_CONFIRM_")
	case model.PaymentTypeTransfert:
		build_and_send_notification(sender.ID, receiver, transaction, `${sender.first_name} ${sender.last_name}  vient de vous faire un transfert d'un montant de ${transaction.amount} cfa`, `${sender.first_name}  ${sender.last_name}  vient de vous envoyer ${transaction.amount} cfa`, transaction.Type.String())
	case model.PaymentTypePaiement:
		build_and_send_notification(sender.ID, receiver, transaction, `${sender.first_name} ${sender.last_name}  vient de vous faire paiement d'un montant de ${transaction.amount} cfa`, `${sender.first_name}  ${sender.last_name}  vient de vous payer ${transaction.amount} cfa`, transaction.Type.String())

	}
}
