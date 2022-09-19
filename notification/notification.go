package notification

import (
	"context"
	"log"
	"math/big"
	"os"
	"server/database"
	"server/finance"
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
	_, err := _collections.UpdateMany(ctx, bson.M{"to": objectId, "isRead": false}, bson.D{{Key: "$set", Value: bson.M{"isRead": true, "updatedAt": _time}}})
	if err != nil {
		return false, err
	}
	return true, nil
}

func SaveNotification(ctx context.Context, notification model.Notification, _to string) {
	_collections := database.MongoClient.Database(os.Getenv("DATABASE")).Collection("notifications")
	from, _ := primitive.ObjectIDFromHex(*notification.FromID)
	to, _ := primitive.ObjectIDFromHex(_to)
	_notification := database.DBNotification{
		ID:        primitive.NewObjectID(),
		Text:      *notification.Text,
		Type:      *notification.Type,
		ImgUrl:    *notification.ImgURL,
		IsRead:    false,
		FromId:    from,
		ToId:      to,
		CreatedAt: *notification.CreatedAt,
	}
	_collections.InsertOne(ctx, _notification)
}

func build_and_send_notification(sender string, receiver model.User, transaction database.DBTransaction, text_database string, text_notif string, _type string) {
	_testing, _ := strconv.ParseBool(os.Getenv("testing"))
	if !_testing {
		snippets.Connect().SendNotificationMEssage(_type, _type, text_notif, transaction.ID.Hex(), *receiver.FcmToken)
	}
}

func CreateDBNotification(_type model.PaymentType, sender model.User, receiver model.User, transaction database.DBTransaction) {
	_amount, _ := new(big.Int).SetString(transaction.Amount, 0)
	amountbigFloat := finance.FromWei(*_amount)
	_time := time.Now().UTC().Format(time.RFC3339)
	switch _type {
	case model.PaymentTypeTopup:
		text := "Votre dépot d'un montant de" + amountbigFloat.String() + "cfa a bien été effectué rendez vous dans votre portefeuille pour le constater"
		notif := model.Notification{
			Text:      &text,
			Type:      &transaction.Type,
			ImgURL:    new(string),
			IsRead:    new(bool),
			FromID:    &sender.ID,
			CreatedAt: &_time,
		}
		SaveNotification(context.Background(), notif, receiver.ID)
		build_and_send_notification(sender.ID, receiver, transaction, text, text, transaction.Type)
	case model.PaymentTypeCommerce:
		text := *sender.FirstName + " " + *sender.LastName + " vient de vous faire paiement d'un montant de" + amountbigFloat.String() + " cfa"
		notif := model.Notification{
			Text:      &text,
			Type:      &transaction.Type,
			ImgURL:    new(string),
			IsRead:    new(bool),
			FromID:    &sender.ID,
			CreatedAt: &_time,
		}
		SaveNotification(context.Background(), notif, receiver.ID)
		build_and_send_notification(sender.ID, receiver, transaction, text, text, transaction.Type)
	case model.PaymentTypeWithdraw:
		text := "Votre de retrait d'un montant de" + amountbigFloat.String() + "cfa a bien été effectué"
		notif := model.Notification{
			Text:      &text,
			Type:      &transaction.Type,
			ImgURL:    new(string),
			IsRead:    new(bool),
			FromID:    &sender.ID,
			CreatedAt: &_time,
		}
		SaveNotification(context.Background(), notif, receiver.ID)
		build_and_send_notification(sender.ID, receiver, transaction, text, text, "WITHDRAW_CONFIRM_")
	case model.PaymentTypeTransfert:
		text := *sender.FirstName + " " + *sender.LastName + " " + "vient de vous faire un transfert d'un montant de " + amountbigFloat.String() + "cfa"
		notif := model.Notification{
			Text:      &text,
			Type:      &transaction.Type,
			ImgURL:    new(string),
			IsRead:    new(bool),
			FromID:    &sender.ID,
			CreatedAt: &_time,
		}
		build_and_send_notification(sender.ID, receiver, transaction, text, text, transaction.Type)
		SaveNotification(context.Background(), notif, receiver.ID)
	case model.PaymentTypePaiement:
		text := *sender.FirstName + " " + *sender.LastName + "  vient de vous faire paiement d'un montant de" + amountbigFloat.String() + " cfa"
		notif := model.Notification{
			Text:      &text,
			Type:      &transaction.Type,
			ImgURL:    new(string),
			IsRead:    new(bool),
			FromID:    &sender.ID,
			CreatedAt: &_time,
		}
		SaveNotification(context.Background(), notif, receiver.ID)
		build_and_send_notification(sender.ID, receiver, transaction, text, text, transaction.Type)

	}
}
