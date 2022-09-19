package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBNotification struct {
	ID        primitive.ObjectID `validate:"nonzero" bson:"_id"`
	Text      string             `validate:"nonzero" bson:"text"`
	Type      string             `validate:"nonzero" bson:"type"`
	ImgUrl    string             `validate:"nonzero" bson:"imgUrl"`
	IsRead    bool               `validate:"nonzero" bson:"isRead"`
	FromId    primitive.ObjectID `validate:"nonzero" bson:"from"`
	ToId      primitive.ObjectID `validate:"nonzero" bson:"to"`
	CreatedAt string             `validate:"nonzero" bson:"createdAt"`
}
