package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBNotification struct {
	Text      string             `validate:"nonzero" bson:"text"`
	Type      string             `validate:"nonzero" bson:"type"`
	ImgUrl    string             `validate:"nonzero" bson:"imgUrl"`
	IsRead    bool               `validate:"nonzero" bson:"isRead"`
	From      primitive.ObjectID `validate:"nonzero" bson:"from"`
	To        primitive.ObjectID `validate:"nonzero" bson:"to"`
	CreatedAt string             `validate:"nonzero" bson:"createdAt"`
}
