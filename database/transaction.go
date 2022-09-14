package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBTransaction struct {
	ID                *primitive.ObjectID `validate:"nonzero" bson:"_id"`
	TransactionId     string              `validate:"nonzero" bson:"transactionId"`
	Source            string              `validate:"nonzero" bson:"source"`
	Destination       string              `validate:"nonzero" bson:"destination"`
	AgencyID          *primitive.ObjectID `validate:"nonzero" bson:"agencyId"`
	ValidatorID       *primitive.ObjectID `validate:"nonzero" bson:"validatorId"`
	CancellorID       *primitive.ObjectID `validate:"nonzero" bson:"cancellorId"`
	Fee               int64               `validate:"nonzero" bson:"fee"`
	FeeEnterprise     int64               `validate:"nonzero" bson:"feeEnterprise"`
	Amount            int64               `validate:"nonzero" bson:"amount"`
	Token             string              `validate:"nonzero" bson:"token"`
	Description       string              `validate:"nonzero" bson:"description"`
	DestinationUserID *primitive.ObjectID `validate:"nonzero" bson:"destinationUserId"`
	ShortId           string              `validate:"nonzero" bson:"shortId"`
	Type              string              `validate:"nonzero" bson:"type"`
	Status            string              `validate:"nonzero" bson:"status"`
	CreatorID         *primitive.ObjectID `validate:"nonzero" bson:"creatorId"`
	EnterpriseID      *primitive.ObjectID `validate:"nonzero" bson:"enterpriseId"`
	Country           string              `validate:"nonzero" bson:"country"`
	CreatedAt         string              `validate:"nonzero" bson:"createdAt"`
	UpdatedAt         string              `validate:"nonzero" bson:"updatedAt"`
	Ip                *ConnectionDB       ` bson:"ip"`
}
